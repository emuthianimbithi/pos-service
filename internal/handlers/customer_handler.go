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

type CustomerHandler struct {
	service *services.CustomerService
}

func NewCustomerHandler(service *services.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

type CreateCustomerRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Email     string `json:"email" binding:"email"`
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	var req CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.service.CreateCustomer(branchID, req.FirstName, req.LastName, req.Phone, req.Email)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Customer created successfully", customer)
}

func (h *CustomerHandler) GetCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	customer, err := h.service.GetCustomer(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Customer not found")
		return
	}

	response.Success(c, http.StatusOK, "Customer retrieved", customer)
}

type UpdateCustomerRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	var req UpdateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	customer, err := h.service.UpdateCustomer(id, req.FirstName, req.LastName, req.Phone, req.Email)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Customer updated successfully", customer)
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid customer ID")
		return
	}

	if err := h.service.DeleteCustomer(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete customer")
		return
	}

	response.Success(c, http.StatusOK, "Customer deleted successfully", nil)
}

func (h *CustomerHandler) ListCustomers(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	customers, total, err := h.service.ListCustomers(branchID, page, limit, search)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch customers")
		return
	}

	response.Success(c, http.StatusOK, "Customers retrieved", gin.H{
		"customers": customers,
		"total":     total,
		"page":      page,
		"limit":     limit,
	})
}
