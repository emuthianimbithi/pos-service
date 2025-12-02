/**
 * Product and Category related types
 */

import { BaseModel } from './common';
import { Branch } from './branch';

export interface Category extends BaseModel {
    branch_id: string;
    name: string;
    description?: string;
    branch?: Branch;
    products?: Product[];
}

export interface Product extends BaseModel {
    branch_id: string;
    category_id?: string;
    name: string;
    description?: string;
    image_url?: string;
    is_active: boolean;
    track_inventory: boolean;
    branch?: Branch;
    category?: Category;
    variants?: ProductVariant[];
}

export interface ProductVariant extends BaseModel {
    product_id: string;
    name: string;
    sku: string;
    price: number;
    cost: number;
    stock: number;
    low_stock_threshold: number;
    product?: Product;
}

export interface CreateCategoryRequest {
    name: string;
    description?: string;
}

export interface UpdateCategoryRequest {
    name?: string;
    description?: string;
}

export interface CreateProductRequest {
    category_id?: string;
    name: string;
    description?: string;
    image_url?: string;
    track_inventory?: boolean;
    variants: CreateProductVariantRequest[];
}

export interface UpdateProductRequest {
    category_id?: string;
    name?: string;
    description?: string;
    image_url?: string;
    is_active?: boolean;
    track_inventory?: boolean;
}

export interface CreateProductVariantRequest {
    name: string;
    sku: string;
    price: number;
    cost: number;
    stock?: number;
    low_stock_threshold?: number;
}

export interface UpdateProductVariantRequest {
    name?: string;
    sku?: string;
    price?: number;
    cost?: number;
    stock?: number;
    low_stock_threshold?: number;
}

export interface ProductListParams {
    category_id?: string;
    is_active?: boolean;
    search?: string;
    track_inventory?: boolean;
}
