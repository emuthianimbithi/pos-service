/**
 * Inventory API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    InventoryTransaction,
    CreateInventoryTransactionRequest,
    InventoryListParams,
    StockAdjustmentRequest,
    LowStockAlert,
} from '../types';

export const inventoryApi = {
    /**
     * Get all inventory transactions
     */
    getTransactions: async (branchId: string, params?: InventoryListParams): Promise<InventoryTransaction[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<InventoryTransaction[]>(`/branches/${branchId}/inventory${query}`);
        return response.data;
    },

    /**
     * Get transaction by ID
     */
    getTransaction: async (branchId: string, id: string): Promise<InventoryTransaction> => {
        const response = await apiClient.get<InventoryTransaction>(`/branches/${branchId}/inventory/${id}`);
        return response.data;
    },

    /**
     * Create inventory transaction
     */
    createTransaction: async (branchId: string, data: CreateInventoryTransactionRequest): Promise<InventoryTransaction> => {
        const response = await apiClient.post<InventoryTransaction>(`/branches/${branchId}/inventory`, data);
        return response.data;
    },

    /**
     * Adjust stock for a variant
     */
    adjustStock: async (branchId: string, data: StockAdjustmentRequest): Promise<InventoryTransaction> => {
        const response = await apiClient.post<InventoryTransaction>(`/branches/${branchId}/inventory/adjust`, data);
        return response.data;
    },

    /**
     * Get low stock alerts
     */
    getLowStockAlerts: async (branchId: string): Promise<LowStockAlert[]> => {
        const response = await apiClient.get<LowStockAlert[]>(`/branches/${branchId}/inventory/low-stock`);
        return response.data;
    },

    /**
     * Get inventory for a specific product
     */
    getProductInventory: async (branchId: string, productId: string): Promise<InventoryTransaction[]> => {
        const response = await apiClient.get<InventoryTransaction[]>(`/branches/${branchId}/inventory/product/${productId}`);
        return response.data;
    },
};
