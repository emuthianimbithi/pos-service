package handlers

import (
	"net/http"
	"strconv"

	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BusinessHandler struct {
	service *services.BusinessService
}

func NewBusinessHandler(service *services.BusinessService) *BusinessHandler {
	return &BusinessHandler{service: service}
}

type CreateBusinessRequest struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	TaxID   string `json:"tax_id"`
}

func (h *BusinessHandler) CreateBusiness(c *gin.Context) {
	var req CreateBusinessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	business, err := h.service.CreateBusiness(req.Name, req.Email, req.Phone, req.Address, req.TaxID)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Business created successfully", business)
}

func (h *BusinessHandler) GetBusiness(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid business ID")
		return
	}

	business, err := h.service.GetBusinessByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Business not found")
		return
	}

	response.Success(c, http.StatusOK, "Business retrieved", business)
}

func (h *BusinessHandler) ListBusinesses(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	businesses, total, err := h.service.ListBusinesses(page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch businesses")
		return
	}

	response.Success(c, http.StatusOK, "Businesses retrieved", gin.H{
		"businesses": businesses,
		"total":      total,
		"page":       page,
		"limit":      limit,
	})
}

type UpdateBusinessRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	TaxID   string `json:"tax_id"`
	LogoURL string `json:"logo_url"`
}

func (h *BusinessHandler) UpdateBusiness(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid business ID")
		return
	}

	var req UpdateBusinessRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	business, err := h.service.UpdateBusiness(id, req.Name, req.Email, req.Phone, req.Address, req.TaxID, req.LogoURL)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Business updated successfully", business)
}

// DeleteBusiness performs soft delete
func (h *BusinessHandler) DeleteBusiness(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid business ID")
		return
	}

	if err := h.service.DeleteBusiness(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete business")
		return
	}

	response.Success(c, http.StatusOK, "Business deactivated successfully", nil)
}
