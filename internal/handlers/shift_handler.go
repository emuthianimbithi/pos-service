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

type ShiftHandler struct {
	service *services.ShiftService
}

func NewShiftHandler(service *services.ShiftService) *ShiftHandler {
	return &ShiftHandler{service: service}
}

type OpenShiftRequest struct {
	OpeningBalance float64 `json:"opening_balance" binding:"required"`
}

func (h *ShiftHandler) OpenShift(c *gin.Context) {
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

	var req OpenShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	shift, err := h.service.OpenShift(branchID, userID, req.OpeningBalance)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Shift opened successfully", shift)
}

type CloseShiftRequest struct {
	ClosingBalance float64 `json:"closing_balance" binding:"required"`
	Notes          string  `json:"notes"`
}

func (h *ShiftHandler) CloseShift(c *gin.Context) {
	userID, err := utils.GetUserID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "User context required")
		return
	}

	var req CloseShiftRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	shift, err := h.service.CloseShift(userID, req.ClosingBalance, req.Notes)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Shift closed successfully", shift)
}

func (h *ShiftHandler) GetShifts(c *gin.Context) {
	userRole, err := utils.GetUserRole(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "User role required")
		return
	}

	var branchID *uuid.UUID
	var userID *uuid.UUID

	// Role-based filtering
	if userRole == models.RoleCashier {
		// Cashiers only see their own shifts
		uid, _ := utils.GetUserID(c)
		userID = &uid
	} else if userRole == models.RoleManager {
		// Managers see all shifts in their branch
		bid, _ := utils.GetBranchID(c)
		branchID = &bid
	} else if userRole == models.RoleAdmin {
		// Admins see all, but can filter by branch if provided in query
		if bidStr := c.Query("branch_id"); bidStr != "" {
			if bid, err := uuid.Parse(bidStr); err == nil {
				branchID = &bid
			}
		}
	}

	shifts, err := h.service.GetShifts(branchID, userID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Shifts retrieved", shifts)
}
