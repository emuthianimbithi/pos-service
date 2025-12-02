/**
 * M-Pesa API functions
 */

import { apiClient } from './client';
import {
    MpesaConfig,
    MpesaPaymentRequest,
    MpesaPaymentResponse,
    UpdateMpesaConfigRequest,
} from '../types';

export const mpesaApi = {
    /**
     * Get M-Pesa configuration
     */
    getConfig: async (): Promise<MpesaConfig> => {
        const response = await apiClient.get<MpesaConfig>('/mpesa/config');
        return response.data;
    },

    /**
     * Update M-Pesa configuration
     */
    updateConfig: async (data: UpdateMpesaConfigRequest): Promise<MpesaConfig> => {
        const response = await apiClient.put<MpesaConfig>('/mpesa/config', data);
        return response.data;
    },

    /**
     * Initiate STK Push payment
     */
    initiatePayment: async (data: MpesaPaymentRequest): Promise<MpesaPaymentResponse> => {
        const response = await apiClient.post<MpesaPaymentResponse>('/mpesa/stk-push', data);
        return response.data;
    },

    /**
     * Query STK Push payment status
     */
    queryPaymentStatus: async (checkoutRequestId: string): Promise<any> => {
        const response = await apiClient.post(`/mpesa/query`, { checkout_request_id: checkoutRequestId });
        return response.data;
    },
};
