/**
 * M-Pesa React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { mpesaApi } from '../api';
import {
    MpesaConfig,
    MpesaPaymentRequest,
    MpesaPaymentResponse,
    UpdateMpesaConfigRequest,
} from '../types';

export const useMpesaConfig = (): UseQueryResult<MpesaConfig, Error> => {
    return useQuery({
        queryKey: ['mpesa', 'config'],
        queryFn: mpesaApi.getConfig,
    });
};

export const useUpdateMpesaConfig = (): UseMutationResult<MpesaConfig, Error, UpdateMpesaConfigRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: mpesaApi.updateConfig,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['mpesa', 'config'] });
        },
    });
};

export const useInitiateMpesaPayment = (): UseMutationResult<MpesaPaymentResponse, Error, MpesaPaymentRequest> => {
    return useMutation({
        mutationFn: mpesaApi.initiatePayment,
    });
};

export const useQueryMpesaPaymentStatus = (checkoutRequestId: string): UseQueryResult<any, Error> => {
    return useQuery({
        queryKey: ['mpesa', 'payment', checkoutRequestId],
        queryFn: () => mpesaApi.queryPaymentStatus(checkoutRequestId),
        enabled: !!checkoutRequestId,
        refetchInterval: 5000, // Poll every 5 seconds
        refetchIntervalInBackground: false,
    });
};
