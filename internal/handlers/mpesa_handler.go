package handlers

import (
	"net/http"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
)

type MpesaHandler struct {
	service *services.MpesaService
}

func NewMpesaHandler(service *services.MpesaService) *MpesaHandler {
	return &MpesaHandler{
		service: service,
	}
}

// CreateConfig creates or updates M-Pesa configuration for a branch
func (h *MpesaHandler) CreateConfig(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		ConsumerKey    string `json:"consumer_key"`
		ConsumerSecret string `json:"consumer_secret"`
		PassKey        string `json:"pass_key"`
		ShortCode      string `json:"short_code"`
		Type           string `json:"type"`
		Environment    string `json:"environment"`
		CallbackURL    string `json:"callback_url"`
	}

	if err := utils.BindJSON(c, &input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	validator := utils.NewValidator()
	validator.Required("consumer_key", input.ConsumerKey)
	validator.Required("consumer_secret", input.ConsumerSecret)
	validator.Required("pass_key", input.PassKey)
	validator.Required("short_code", input.ShortCode)

	if !validator.IsValid() {
		response.ValidationError(c, validator.Errors().ToMap())
		return
	}

	config := &models.MpesaConfig{
		BranchModel: models.BranchModel{
			TenantModel: models.TenantModel{
				BusinessID: businessID,
			},
			BranchID: branchID,
		},
		ConsumerKey:    input.ConsumerKey,
		ConsumerSecret: input.ConsumerSecret,
		PassKey:        input.PassKey,
		ShortCode:      input.ShortCode,
		Type:           input.Type,
		Environment:    input.Environment,
		CallbackURL:    input.CallbackURL,
	}

	// Save configuration via service
	if err := h.service.SaveConfig(config); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to save configuration")
		return
	}

	response.Success(c, http.StatusOK, "Configuration saved successfully", config)
}

// InitiateSTKPush triggers an STK Push
func (h *MpesaHandler) InitiateSTKPush(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	var input struct {
		Amount      float64 `json:"amount"`
		PhoneNumber string  `json:"phone_number"`
		Reference   string  `json:"reference"`
		Description string  `json:"description"`
	}

	if err := utils.BindJSON(c, &input); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	validator := utils.NewValidator()
	validator.Required("phone_number", input.PhoneNumber)
	validator.Min("amount", input.Amount, 1)

	if !validator.IsValid() {
		response.ValidationError(c, validator.Errors().ToMap())
		return
	}

	tx, err := h.service.InitiateSTKPush(branchID, input.Amount, input.PhoneNumber, input.Reference, input.Description)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "STK Push initiated successfully", tx)
}

// Callback handles Safaricom callbacks
func (h *MpesaHandler) Callback(c *gin.Context) {
	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}

	if err := h.service.ProcessCallback(body); err != nil {
		// Log error but return 200 to Safaricom so they don't retry
		// log.Printf("Failed to process callback: %v", err)
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// CheckStatus checks the status of a transaction
func (h *MpesaHandler) CheckStatus(c *gin.Context) {
	id, err := utils.ParseUUIDParam(c, "id")
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	tx, err := h.service.GetTransactionStatus(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Transaction not found")
		return
	}

	response.Success(c, http.StatusOK, "Transaction status retrieved", gin.H{
		"id":          tx.ID,
		"status":      tx.Status,
		"result_code": tx.ResultCode,
		"result_desc": tx.ResultDesc,
		"amount":      tx.Amount,
		"reference":   tx.Reference,
	})
}
