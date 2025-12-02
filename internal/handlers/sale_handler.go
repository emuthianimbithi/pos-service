package handlers

import (
	"net/http"

	"github.com/emuthianimbithi/pos-service/internal/models"
	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SaleHandler struct {
	service *services.SaleService
}

func NewSaleHandler(service *services.SaleService) *SaleHandler {
	return &SaleHandler{service: service}
}

type CreateSaleRequest struct {
	Items []struct {
		ProductID uuid.UUID  `json:"product_id" binding:"required"`
		VariantID *uuid.UUID `json:"variant_id"`
		Quantity  int        `json:"quantity" binding:"required,min=1"`
	} `json:"items" binding:"required,dive"`
	PaymentMethod string `json:"payment_method" binding:"required,oneof=cash card mpesa"`
	Notes         string `json:"notes"`
}

func (h *SaleHandler) CreateSale(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "User context required")
		return
	}

	var req CreateSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// Map request items to service input
	var serviceItems []services.SaleItemInput
	for _, item := range req.Items {
		serviceItems = append(serviceItems, services.SaleItemInput{
			ProductID: item.ProductID,
			VariantID: item.VariantID,
			Quantity:  item.Quantity,
		})
	}

	sale, err := h.service.ProcessSale(branchID, userID, serviceItems, req.PaymentMethod, req.Notes)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Sale processed successfully", sale)
}

func (h *SaleHandler) GetSales(c *gin.Context) {
	userRole, err := utils.GetUserRole(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "User role required")
		return
	}

	var branchID *uuid.UUID
	var userID *uuid.UUID

	// Role-based filtering
	if userRole == models.RoleCashier {
		uid, _ := utils.GetUserID(c)
		userID = &uid
	} else if userRole == models.RoleManager {
		bid, _ := utils.GetBranchID(c)
		branchID = &bid
	} else if userRole == models.RoleAdmin {
		if bidStr := c.Query("branch_id"); bidStr != "" {
			if bid, err := uuid.Parse(bidStr); err == nil {
				branchID = &bid
			}
		}
	}

	sales, err := h.service.GetSales(branchID, userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Sales retrieved", sales)
}

func (h *SaleHandler) GetSaleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid sale ID")
		return
	}

	sale, err := h.service.GetSaleByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Sale not found")
		return
	}

	response.Success(c, http.StatusOK, "Sale details retrieved", sale)
}

type RefundSaleRequest struct {
	Reason string `json:"reason" binding:"required"`
}

func (h *SaleHandler) RefundSale(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid sale ID")
		return
	}

	userID, err := utils.GetUserID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "User context required")
		return
	}

	var req RefundSaleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	sale, err := h.service.RefundSale(id, userID, req.Reason)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Sale refunded successfully", sale)
}
