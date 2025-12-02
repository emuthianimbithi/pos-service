/**
 * Product API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    Product,
    ProductVariant,
    Category,
    CreateProductRequest,
    UpdateProductRequest,
    CreateProductVariantRequest,
    UpdateProductVariantRequest,
    CreateCategoryRequest,
    UpdateCategoryRequest,
    ProductListParams,
} from '../types';

export const productApi = {
    // Categories
    /**
     * Get all categories for current branch
     */
    getCategories: async (branchId: string): Promise<Category[]> => {
        const response = await apiClient.get<Category[]>(`/branches/${branchId}/categories`);
        return response.data;
    },

    /**
     * Get category by ID
     */
    getCategory: async (branchId: string, id: string): Promise<Category> => {
        const response = await apiClient.get<Category>(`/branches/${branchId}/categories/${id}`);
        return response.data;
    },

    /**
     * Create a new category
     */
    createCategory: async (branchId: string, data: CreateCategoryRequest): Promise<Category> => {
        const response = await apiClient.post<Category>(`/branches/${branchId}/categories`, data);
        return response.data;
    },

    /**
     * Update category by ID
     */
    updateCategory: async (branchId: string, id: string, data: UpdateCategoryRequest): Promise<Category> => {
        const response = await apiClient.put<Category>(`/branches/${branchId}/categories/${id}`, data);
        return response.data;
    },

    /**
     * Delete category by ID
     */
    deleteCategory: async (branchId: string, id: string): Promise<void> => {
        await apiClient.delete(`/branches/${branchId}/categories/${id}`);
    },

    // Products
    /**
     * Get all products with optional filters
     */
    getProducts: async (branchId: string, params?: ProductListParams): Promise<Product[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Product[]>(`/branches/${branchId}/products${query}`);
        return response.data;
    },

    /**
     * Get product by ID
     */
    getProduct: async (branchId: string, id: string): Promise<Product> => {
        const response = await apiClient.get<Product>(`/branches/${branchId}/products/${id}`);
        return response.data;
    },

    /**
     * Create a new product with variants
     */
    createProduct: async (branchId: string, data: CreateProductRequest): Promise<Product> => {
        const response = await apiClient.post<Product>(`/branches/${branchId}/products`, data);
        return response.data;
    },

    /**
     * Update product by ID
     */
    updateProduct: async (branchId: string, id: string, data: UpdateProductRequest): Promise<Product> => {
        const response = await apiClient.put<Product>(`/branches/${branchId}/products/${id}`, data);
        return response.data;
    },

    /**
     * Delete product by ID
     */
    deleteProduct: async (branchId: string, id: string): Promise<void> => {
        await apiClient.delete(`/branches/${branchId}/products/${id}`);
    },

    // Product Variants
    /**
     * Get all variants for a product
     */
    getProductVariants: async (branchId: string, productId: string): Promise<ProductVariant[]> => {
        const response = await apiClient.get<ProductVariant[]>(`/branches/${branchId}/products/${productId}/variants`);
        return response.data;
    },

    /**
     * Get variant by ID
     */
    getProductVariant: async (branchId: string, productId: string, variantId: string): Promise<ProductVariant> => {
        const response = await apiClient.get<ProductVariant>(`/branches/${branchId}/products/${productId}/variants/${variantId}`);
        return response.data;
    },

    /**
     * Create a new product variant
     */
    createProductVariant: async (branchId: string, productId: string, data: CreateProductVariantRequest): Promise<ProductVariant> => {
        const response = await apiClient.post<ProductVariant>(`/branches/${branchId}/products/${productId}/variants`, data);
        return response.data;
    },

    /**
     * Update product variant
     */
    updateProductVariant: async (branchId: string, productId: string, variantId: string, data: UpdateProductVariantRequest): Promise<ProductVariant> => {
        const response = await apiClient.put<ProductVariant>(`/branches/${branchId}/products/${productId}/variants/${variantId}`, data);
        return response.data;
    },

    /**
     * Delete product variant
     */
    deleteProductVariant: async (branchId: string, productId: string, variantId: string): Promise<void> => {
        await apiClient.delete(`/branches/${branchId}/products/${productId}/variants/${variantId}`);
    },
};
