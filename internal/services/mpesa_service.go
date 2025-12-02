package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/repository"
	"github.com/google/uuid"
)

const (
	SandboxURL    = "https://sandbox.safaricom.co.ke"
	ProductionURL = "https://api.safaricom.co.ke"
)

type MpesaService struct {
	repo *repository.MpesaRepository
}

func NewMpesaService(repo *repository.MpesaRepository) *MpesaService {
	return &MpesaService{repo: repo}
}

// STKPushRequest represents the payload for STK Push
type STKPushRequest struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// SaveConfig saves or updates M-Pesa configuration
func (s *MpesaService) SaveConfig(config *models.MpesaConfig) error {
	return s.repo.UpsertConfig(config)
}

// InitiateSTKPush starts an M-Pesa transaction
func (s *MpesaService) InitiateSTKPush(branchID uuid.UUID, amount float64, phoneNumber, reference, description string) (*models.MpesaTransaction, error) {
	// 1. Get Branch Configuration
	config, err := s.repo.GetConfigByBranchID(branchID)
	if err != nil {
		return nil, errors.New("mpesa configuration not found for this branch")
	}

	// 2. Generate Access Token
	token, err := s.getAccessToken(config)
	if err != nil {
		return nil, err
	}

	// 3. Prepare Request
	timestamp := time.Now().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(config.ShortCode + config.PassKey + timestamp))

	baseURL := SandboxURL
	if config.Environment == "production" {
		baseURL = ProductionURL
	}

	// Determine transaction type
	txType := "CustomerPayBillOnline"
	if config.Type == "till_number" {
		txType = "CustomerBuyGoodsOnline"
	}

	reqBody := STKPushRequest{
		BusinessShortCode: config.ShortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   txType,
		Amount:            fmt.Sprintf("%.0f", amount),
		PartyA:            phoneNumber,
		PartyB:            config.ShortCode,
		PhoneNumber:       phoneNumber,
		CallBackURL:       config.CallbackURL, // Should be your API endpoint
		AccountReference:  reference,
		TransactionDesc:   description,
	}

	jsonBody, _ := json.Marshal(reqBody)

	// 4. Send Request
	req, _ := http.NewRequest("POST", baseURL+"/mpesa/stkpush/v1/processrequest", bytes.NewBuffer(jsonBody))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("mpesa api error: %s", string(body))
	}

	var mpesaResp struct {
		MerchantRequestID   string `json:"MerchantRequestID"`
		CheckoutRequestID   string `json:"CheckoutRequestID"`
		ResponseCode        string `json:"ResponseCode"`
		ResponseDescription string `json:"ResponseDescription"`
		CustomerMessage     string `json:"CustomerMessage"`
	}

	if err := json.Unmarshal(body, &mpesaResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	// 5. Save Transaction Record
	tx := &models.MpesaTransaction{
		BranchModel: models.BranchModel{
			TenantModel: models.TenantModel{
				BaseModel:  models.BaseModel{ID: uuid.New()},
				BusinessID: config.BusinessID,
			},
			BranchID: branchID,
		},
		MerchantRequestID: mpesaResp.MerchantRequestID,
		CheckoutRequestID: mpesaResp.CheckoutRequestID,
		Amount:            amount,
		PhoneNumber:       phoneNumber,
		Reference:         reference,
		Description:       description,
		Status:            "pending",
		ResultDesc:        mpesaResp.ResponseDescription,
	}

	if err := s.repo.CreateTransaction(tx); err != nil {
		return nil, err
	}

	return tx, nil
}

// MpesaCallback represents the callback payload from Safaricom
type MpesaCallback struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

// ProcessCallback processes the M-Pesa callback
func (s *MpesaService) ProcessCallback(payload []byte) error {
	var callback MpesaCallback
	if err := json.Unmarshal(payload, &callback); err != nil {
		return err
	}

	data := callback.Body.StkCallback

	// Find transaction
	tx, err := s.repo.GetTransactionByCheckoutRequestID(data.CheckoutRequestID)
	if err != nil {
		return errors.New("transaction not found")
	}

	// Determine status
	status := "failed"
	if data.ResultCode == 0 {
		status = "completed"
	} else if data.ResultCode == 1032 {
		status = "cancelled"
	}

	// Extract receipt number if successful
	var receipt string
	if status == "completed" {
		for _, item := range data.CallbackMetadata.Item {
			if item.Name == "MpesaReceiptNumber" {
				if val, ok := item.Value.(string); ok {
					receipt = val
				}
				break
			}
		}
	}

	// Update transaction
	// We use a custom update map here to include the receipt number
	updates := map[string]interface{}{
		"status":      status,
		"result_code": data.ResultCode,
		"result_desc": data.ResultDesc,
	}

	if receipt != "" {
		updates["mpesa_receipt"] = receipt
	}

	// Use the repository to update
	return s.repo.UpdateTransactionStatus(tx.ID, status, data.ResultCode, data.ResultDesc, receipt)
}

// GetTransactionStatus retrieves the status of a transaction
func (s *MpesaService) GetTransactionStatus(id uuid.UUID) (*models.MpesaTransaction, error) {
	return s.repo.GetTransactionByID(id)
}

func (s *MpesaService) getAccessToken(config *models.MpesaConfig) (string, error) {
	baseURL := SandboxURL
	if config.Environment == "production" {
		baseURL = ProductionURL
	}

	req, _ := http.NewRequest("GET", baseURL+"/oauth/v1/generate?grant_type=client_credentials", nil)
	auth := base64.StdEncoding.EncodeToString([]byte(config.ConsumerKey + ":" + config.ConsumerSecret))
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get access token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("failed to authenticate with mpesa")
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   string `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}
