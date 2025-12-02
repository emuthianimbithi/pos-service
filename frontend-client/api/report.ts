/**
 * Report API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    SalesReport,
    InventoryReport,
    RevenueReport,
    CustomerReport,
    ReportParams,
} from '../types';

export const reportApi = {
    /**
     * Get sales report
     */
    getSalesReport: async (branchId: string, params: ReportParams): Promise<SalesReport> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<SalesReport>(`/branches/${branchId}/reports/sales${query}`);
        return response.data;
    },

    /**
     * Get inventory report
     */
    getInventoryReport: async (branchId: string): Promise<InventoryReport> => {
        const response = await apiClient.get<InventoryReport>(`/branches/${branchId}/reports/inventory`);
        return response.data;
    },

    /**
     * Get revenue/profit report
     */
    getRevenueReport: async (branchId: string, params: ReportParams): Promise<RevenueReport> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<RevenueReport>(`/branches/${branchId}/reports/revenue${query}`);
        return response.data;
    },

    /**
     * Get customer report
     */
    getCustomerReport: async (branchId: string, params?: Partial<ReportParams>): Promise<CustomerReport> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<CustomerReport>(`/branches/${branchId}/reports/customers${query}`);
        return response.data;
    },

    /**
     * Export sales report as CSV
     */
    exportSalesReportCSV: async (branchId: string, params: ReportParams): Promise<Blob> => {
        const query = buildQueryString(params);
        const response = await apiClient.get(`/branches/${branchId}/reports/sales/export${query}`, {
            responseType: 'blob',
        });
        return response.data;
    },

    /**
     * Export inventory report as CSV
     */
    exportInventoryReportCSV: async (branchId: string): Promise<Blob> => {
        const response = await apiClient.get(`/branches/${branchId}/reports/inventory/export`, {
            responseType: 'blob',
        });
        return response.data;
    },
};
