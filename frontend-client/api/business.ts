/**
 * Business API functions
 */

import { apiClient } from './client';
import {
    Business,
    CreateBusinessRequest,
    UpdateBusinessRequest,
} from '../types';

export const businessApi = {
    /**
     * Get current business
     */
    getBusiness: async (): Promise<Business> => {
        const response = await apiClient.get<Business>('/business');
        return response.data;
    },

    /**
     * Update current business
     */
    updateBusiness: async (data: UpdateBusinessRequest): Promise<Business> => {
        const response = await apiClient.put<Business>('/business', data);
        return response.data;
    },

    /**
     * Upload business logo
     */
    uploadLogo: async (file: File): Promise<{ logo_url: string }> => {
        const formData = new FormData();
        formData.append('logo', file);

        const response = await apiClient.post<{ logo_url: string }>(
            '/business/logo',
            formData,
            {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            }
        );
        return response.data;
    },
};
