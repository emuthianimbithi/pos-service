package handlers

import (
	"net/http"

	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type InventoryHandler struct {
	service *services.InventoryService
}

func NewInventoryHandler(service *services.InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

type AdjustStockRequest struct {
	VariantID uuid.UUID `json:"variant_id" binding:"required"`
	Quantity  int       `json:"quantity" binding:"required"` // Can be negative
	Reason    string    `json:"reason" binding:"required"`
}

func (h *InventoryHandler) AdjustStock(c *gin.Context) {
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

	var req AdjustStockRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.AdjustStock(branchID, req.VariantID, userID, req.Quantity, req.Reason); err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Stock adjusted successfully", nil)
}
