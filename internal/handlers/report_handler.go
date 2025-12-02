package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GetSalesReport(c *gin.Context) {
	var branchID *uuid.UUID

	// Try to get branch from context (for managers/cashiers)
	if bid, err := utils.GetBranchID(c); err == nil {
		branchID = &bid
	}

	// Allow admin to specify branch_id
	if bidStr := c.Query("branch_id"); bidStr != "" {
		if bid, err := uuid.Parse(bidStr); err == nil {
			branchID = &bid
		}
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		response.Error(c, http.StatusBadRequest, "start_date and end_date are required (format: 2006-01-02)")
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid start_date format (use: 2006-01-02)")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid end_date format (use: 2006-01-02)")
		return
	}

	// Set end date to end of day
	endDate = endDate.Add(23*time.Hour + 59*time.Minute + 59*time.Second)

	report, err := h.service.GetSalesReport(branchID, startDate, endDate)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate sales report")
		return
	}

	response.Success(c, http.StatusOK, "Sales report generated", report)
}

func (h *ReportHandler) GetInventoryReport(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	report, err := h.service.GetInventoryReport(branchID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate inventory report")
		return
	}

	response.Success(c, http.StatusOK, "Inventory report generated", report)
}

func (h *ReportHandler) GetLowStockReport(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	products, err := h.service.GetLowStockReport(branchID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate low stock report")
		return
	}

	response.Success(c, http.StatusOK, "Low stock report generated", products)
}

func (h *ReportHandler) GetCustomerReport(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	report, err := h.service.GetCustomerReport(branchID, limit)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to generate customer report")
		return
	}

	response.Success(c, http.StatusOK, "Customer report generated", report)
}
