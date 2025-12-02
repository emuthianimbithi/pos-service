/**
 * Sale API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    Sale,
    CreateSaleRequest,
    SaleListParams,
} from '../types';

export const saleApi = {
    /**
     * Get all sales with optional filters
     */
    getSales: async (branchId: string, params?: SaleListParams): Promise<Sale[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Sale[]>(`/branches/${branchId}/sales${query}`);
        return response.data;
    },

    /**
     * Get sale by ID
     */
    getSale: async (branchId: string, id: string): Promise<Sale> => {
        const response = await apiClient.get<Sale>(`/branches/${branchId}/sales/${id}`);
        return response.data;
    },

    /**
     * Create a new sale
     */
    createSale: async (branchId: string, data: CreateSaleRequest): Promise<Sale> => {
        const response = await apiClient.post<Sale>(`/branches/${branchId}/sales`, data);
        return response.data;
    },

    /**
     * Get sale by receipt number
     */
    getSaleByReceipt: async (branchId: string, receiptNumber: string): Promise<Sale> => {
        const response = await apiClient.get<Sale>(`/branches/${branchId}/sales/receipt/${receiptNumber}`);
        return response.data;
    },

    /**
     * Refund a sale
     */
    refundSale: async (branchId: string, id: string, reason?: string): Promise<Sale> => {
        const response = await apiClient.post<Sale>(`/branches/${branchId}/sales/${id}/refund`, { reason });
        return response.data;
    },
};
