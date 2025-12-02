/**
 * Customer API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    Customer,
    CreateCustomerRequest,
    UpdateCustomerRequest,
    CustomerListParams,
} from '../types';

export const customerApi = {
    /**
     * Get all customers with optional filters
     */
    getCustomers: async (branchId: string, params?: CustomerListParams): Promise<Customer[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Customer[]>(`/branches/${branchId}/customers${query}`);
        return response.data;
    },

    /**
     * Get customer by ID
     */
    getCustomer: async (branchId: string, id: string): Promise<Customer> => {
        const response = await apiClient.get<Customer>(`/branches/${branchId}/customers/${id}`);
        return response.data;
    },

    /**
     * Create a new customer
     */
    createCustomer: async (branchId: string, data: CreateCustomerRequest): Promise<Customer> => {
        const response = await apiClient.post<Customer>(`/branches/${branchId}/customers`, data);
        return response.data;
    },

    /**
     * Update customer by ID
     */
    updateCustomer: async (branchId: string, id: string, data: UpdateCustomerRequest): Promise<Customer> => {
        const response = await apiClient.put<Customer>(`/branches/${branchId}/customers/${id}`, data);
        return response.data;
    },

    /**
     * Delete customer by ID
     */
    deleteCustomer: async (branchId: string, id: string): Promise<void> => {
        await apiClient.delete(`/branches/${branchId}/customers/${id}`);
    },

    /**
     * Search customers by phone number
     */
    searchByPhone: async (branchId: string, phone: string): Promise<Customer[]> => {
        const response = await apiClient.get<Customer[]>(`/branches/${branchId}/customers/search?phone=${phone}`);
        return response.data;
    },
};
