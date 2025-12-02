/**
 * Inventory related types
 */

import { BaseModel, TransactionType, StockStatus } from './common';
import { Branch } from './branch';
import { ProductVariant } from './product';
import { User } from './user';

export interface InventoryTransaction extends BaseModel {
    branch_id: string;
    product_variant_id: string;
    user_id: string;
    type: TransactionType;
    quantity: number;
    reason?: string;
    reference_id?: string;
    branch?: Branch;
    variant?: ProductVariant;
    user?: User;
}

export interface CreateInventoryTransactionRequest {
    product_variant_id: string;
    type: TransactionType;
    quantity: number;
    reason?: string;
}

export interface InventoryListParams {
    product_id?: string;
    variant_id?: string;
    type?: TransactionType;
    start_date?: string;
    end_date?: string;
}

export interface StockAdjustmentRequest {
    variant_id: string;
    quantity: number;
    reason: string;
}

export interface LowStockAlert {
    variant_id: string;
    variant_name: string;
    sku: string;
    current_stock: number;
    low_stock_threshold: number;
    product_name: string;
}
