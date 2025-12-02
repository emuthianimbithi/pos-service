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

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	role := c.Query("role")

	var branchID *uuid.UUID
	if branchIDStr := c.Query("branch_id"); branchIDStr != "" {
		if bid, err := uuid.Parse(branchIDStr); err == nil {
			branchID = &bid
		}
	}

	users, total, err := h.service.ListUsers(businessID, branchID, role, page, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	response.Success(c, http.StatusOK, "Users retrieved", gin.H{
		"users": users,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

type CreateUserRequest struct {
	Email     string     `json:"email" binding:"required,email"`
	Password  string     `json:"password" binding:"required,min=6"`
	FirstName string     `json:"first_name" binding:"required"`
	LastName  string     `json:"last_name" binding:"required"`
	Phone     string     `json:"phone" binding:"required"`
	Role      string     `json:"role" binding:"required,oneof=admin manager cashier"`
	BranchID  *uuid.UUID `json:"branch_id"`
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	businessID, err := utils.GetBusinessID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Business context required")
		return
	}

	var req CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.CreateUser(businessID, req.BranchID, req.Email, req.Password, req.FirstName, req.LastName, req.Phone, req.Role)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", user)
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := h.service.GetUserProfile(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.Success(c, http.StatusOK, "User profile retrieved", user)
}

type UpdateProfileRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.service.UpdateUserProfile(id, req.FirstName, req.LastName, req.Phone)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User profile updated successfully", user)
}

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

func (h *UserHandler) ChangePassword(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.ChangePassword(id, req.NewPassword); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Password changed successfully", nil)
}

func (h *UserHandler) DeactivateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.service.DeactivateUser(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to deactivate user")
		return
	}

	response.Success(c, http.StatusOK, "User deactivated successfully", nil)
}
