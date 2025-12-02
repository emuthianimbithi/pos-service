/**
 * Sale related types
 */

import { BaseModel, PaymentMethod, SaleStatus } from './common';
import { Branch } from './branch';
import { Customer } from './customer';
import { User } from './user';
import { Product } from './product';

export interface Sale extends BaseModel {
    branch_id: string;
    customer_id?: string;
    user_id: string;
    shift_id?: string;
    receipt_number: string;
    subtotal: number;
    tax: number;
    discount: number;
    total: number;
    payment_method: PaymentMethod;
    status: SaleStatus;
    notes?: string;
    branch?: Branch;
    customer?: Customer;
    user?: User;
    sale_items?: SaleItem[];
}

export interface SaleItem extends BaseModel {
    sale_id: string;
    product_id: string;
    variant_id?: string;
    quantity: number;
    unit_price: number;
    subtotal: number;
    discount: number;
    total: number;
    sale?: Sale;
    product?: Product;
}

export interface CreateSaleRequest {
    customer_id?: string;
    shift_id?: string;
    payment_method: PaymentMethod;
    tax?: number;
    discount?: number;
    notes?: string;
    items: CreateSaleItemRequest[];
}

export interface CreateSaleItemRequest {
    product_id: string;
    variant_id?: string;
    quantity: number;
    unit_price: number;
    discount?: number;
}

export interface SaleListParams {
    customer_id?: string;
    user_id?: string;
    shift_id?: string;
    payment_method?: PaymentMethod;
    status?: SaleStatus;
    start_date?: string;
    end_date?: string;
    search?: string;
}
