package handlers

import (
	"net/http"
	"strconv"

	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BranchHandler struct {
	service *services.BranchService
}

func NewBranchHandler(service *services.BranchService) *BranchHandler {
	return &BranchHandler{service: service}
}

type CreateBranchRequest struct {
	Name    string `json:"name" binding:"required"`
	Code    string `json:"code" binding:"required"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (h *BranchHandler) CreateBranch(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	var req CreateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	branch, err := h.service.CreateBranch(businessID, req.Name, req.Code, req.Phone, req.Address)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Branch created successfully", branch)
}

func (h *BranchHandler) GetBranch(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid branch ID")
		return
	}

	branch, err := h.service.GetBranchByID(id, businessID)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Branch not found")
		return
	}

	response.Success(c, http.StatusOK, "Branch retrieved", branch)
}

func (h *BranchHandler) ListBranches(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	branches, total, err := h.service.ListBranches(businessID, page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch branches")
		return
	}

	response.Success(c, http.StatusOK, "Branches retrieved", gin.H{
		"branches": branches,
		"total":    total,
		"page":     page,
		"limit":    limit,
	})
}

type UpdateBranchRequest struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (h *BranchHandler) UpdateBranch(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid branch ID")
		return
	}

	var req UpdateBranchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	branch, err := h.service.UpdateBranch(id, businessID, req.Name, req.Code, req.Phone, req.Address)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Branch updated successfully", branch)
}

// DeleteBranch performs soft delete
func (h *BranchHandler) DeleteBranch(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid branch ID")
		return
	}

	if err := h.service.DeleteBranch(id, businessID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete branch")
		return
	}

	response.Success(c, http.StatusOK, "Branch deactivated successfully", nil)
}
