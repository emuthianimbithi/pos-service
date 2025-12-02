/**
 * Shift API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    Shift,
    CreateShiftRequest,
    CloseShiftRequest,
    ShiftListParams,
} from '../types';

export const shiftApi = {
    /**
     * Get all shifts with optional filters
     */
    getShifts: async (branchId: string, params?: ShiftListParams): Promise<Shift[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Shift[]>(`/branches/${branchId}/shifts${query}`);
        return response.data;
    },

    /**
     * Get shift by ID
     */
    getShift: async (branchId: string, id: string): Promise<Shift> => {
        const response = await apiClient.get<Shift>(`/branches/${branchId}/shifts/${id}`);
        return response.data;
    },

    /**
     * Start a new shift
     */
    startShift: async (branchId: string, data: CreateShiftRequest): Promise<Shift> => {
        const response = await apiClient.post<Shift>(`/branches/${branchId}/shifts`, data);
        return response.data;
    },

    /**
     * Close a shift
     */
    closeShift: async (branchId: string, id: string, data: CloseShiftRequest): Promise<Shift> => {
        const response = await apiClient.post<Shift>(`/branches/${branchId}/shifts/${id}/close`, data);
        return response.data;
    },

    /**
     * Get current active shift for user
     */
    getActiveShift: async (branchId: string): Promise<Shift | null> => {
        const response = await apiClient.get<Shift | null>(`/branches/${branchId}/shifts/active`);
        return response.data;
    },

    /**
     * Get shift summary/statistics
     */
    getShiftSummary: async (branchId: string, id: string): Promise<Shift> => {
        const response = await apiClient.get<Shift>(`/branches/${branchId}/shifts/${id}/summary`);
        return response.data;
    },
};
