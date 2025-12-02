/**
 * Branch API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    Branch,
    CreateBranchRequest,
    UpdateBranchRequest,
    BranchListParams,
} from '../types';

export const branchApi = {
    /**
     * Get all branches with optional filters
     */
    getBranches: async (params?: BranchListParams): Promise<Branch[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Branch[]>(`/branches${query}`);
        return response.data;
    },

    /**
     * Get branch by ID
     */
    getBranch: async (id: string): Promise<Branch> => {
        const response = await apiClient.get<Branch>(`/branches/${id}`);
        return response.data;
    },

    /**
     * Create a new branch
     */
    createBranch: async (data: CreateBranchRequest): Promise<Branch> => {
        const response = await apiClient.post<Branch>('/branches', data);
        return response.data;
    },

    /**
     * Update branch by ID
     */
    updateBranch: async (id: string, data: UpdateBranchRequest): Promise<Branch> => {
        const response = await apiClient.put<Branch>(`/branches/${id}`, data);
        return response.data;
    },

    /**
     * Delete branch by ID
     */
    deleteBranch: async (id: string): Promise<void> => {
        await apiClient.delete(`/branches/${id}`);
    },
};
