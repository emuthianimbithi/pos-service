package handlers

import (
	"net/http"

	"github.com/emuthianimbithi/pos-service/internal/services"
	"github.com/emuthianimbithi/pos-service/internal/utils"
	"github.com/emuthianimbithi/pos-service/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// Categories
type CreateCategoryRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (h *ProductHandler) CreateCategory(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	var req CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	category, err := h.service.CreateCategory(branchID, req.Name, req.Description)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Category created", category)
}

func (h *ProductHandler) GetCategories(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	categories, err := h.service.GetCategories(branchID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Categories retrieved", categories)
}

// Products
type CreateProductRequest struct {
	CategoryID     *uuid.UUID `json:"category_id"`
	Name           string     `json:"name" binding:"required"`
	Description    string     `json:"description"`
	ImageURL       string     `json:"image_url"`
	TrackInventory bool       `json:"track_inventory"`
	Variants       []struct {
		Name              string  `json:"name" binding:"required"`
		SKU               string  `json:"sku" binding:"required"`
		Price             float64 `json:"price" binding:"required"`
		Cost              float64 `json:"cost"`
		Stock             int     `json:"stock"`
		LowStockThreshold int     `json:"low_stock_threshold"`
	} `json:"variants"`
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	var req CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Create Product
	product, err := h.service.CreateProduct(branchID, req.CategoryID, req.Name, req.Description, req.ImageURL, req.TrackInventory)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 2. Create Variants
	for _, v := range req.Variants {
		_, err := h.service.AddVariant(product.ID, v.Name, v.SKU, v.Price, v.Cost, v.Stock, v.LowStockThreshold)
		if err != nil {
			// Note: In a real app, we should rollback the product creation here (transaction).
			// For now, we just return the error.
			response.Error(c, http.StatusInternalServerError, "Failed to create variant: "+err.Error())
			return
		}
	}

	response.Success(c, http.StatusCreated, "Product created", product)
}

type CreateVariantRequest struct {
	Name              string  `json:"name" binding:"required"`
	SKU               string  `json:"sku" binding:"required"`
	Price             float64 `json:"price" binding:"required"`
	Cost              float64 `json:"cost"`
	Stock             int     `json:"stock"`
	LowStockThreshold int     `json:"low_stock_threshold"`
}

func (h *ProductHandler) CreateVariant(c *gin.Context) {
	productIDStr := c.Param("id")
	productID, err := uuid.Parse(productIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req CreateVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	variant, err := h.service.AddVariant(productID, req.Name, req.SKU, req.Price, req.Cost, req.Stock, req.LowStockThreshold)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "Variant created", variant)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	branchID, err := utils.GetBranchID(c)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Branch context required")
		return
	}

	products, err := h.service.GetProducts(branchID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Products retrieved", products)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	product, err := h.service.GetProduct(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Product not found")
		return
	}

	response.Success(c, http.StatusOK, "Product retrieved", product)
}

type UpdateProductRequest struct {
	CategoryID     *uuid.UUID `json:"category_id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	ImageURL       string     `json:"image_url"`
	TrackInventory bool       `json:"track_inventory"`
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	var req UpdateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.service.UpdateProduct(id, req.CategoryID, req.Name, req.Description, req.ImageURL, req.TrackInventory)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Product updated successfully", product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid product ID")
		return
	}

	if err := h.service.DeleteProduct(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete product")
		return
	}

	response.Success(c, http.StatusOK, "Product deactivated successfully", nil)
}

type UpdateVariantRequest struct {
	Name              string  `json:"name"`
	SKU               string  `json:"sku"`
	Price             float64 `json:"price"`
	Cost              float64 `json:"cost"`
	Stock             int     `json:"stock"`
	LowStockThreshold int     `json:"low_stock_threshold"`
}

func (h *ProductHandler) UpdateVariant(c *gin.Context) {
	variantIDStr := c.Param("variant_id")
	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	var req UpdateVariantRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	variant, err := h.service.UpdateVariant(variantID, req.Name, req.SKU, req.Price, req.Cost, req.Stock, req.LowStockThreshold)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	response.Success(c, http.StatusOK, "Variant updated successfully", variant)
}

func (h *ProductHandler) DeleteVariant(c *gin.Context) {
	variantIDStr := c.Param("variant_id")
	variantID, err := uuid.Parse(variantIDStr)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid variant ID")
		return
	}

	if err := h.service.DeleteVariant(variantID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete variant")
		return
	}

	response.Success(c, http.StatusOK, "Variant deleted successfully", nil)
}
