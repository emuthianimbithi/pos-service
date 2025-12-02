/**
 * Sale React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { saleApi } from '../api';
import {
    Sale,
    CreateSaleRequest,
    SaleListParams,
} from '../types';

export const useSales = (branchId: string, params?: SaleListParams): UseQueryResult<Sale[], Error> => {
    return useQuery({
        queryKey: ['sales', branchId, params],
        queryFn: () => saleApi.getSales(branchId, params),
        enabled: !!branchId,
    });
};

export const useSale = (branchId: string, id: string): UseQueryResult<Sale, Error> => {
    return useQuery({
        queryKey: ['sales', branchId, id],
        queryFn: () => saleApi.getSale(branchId, id),
        enabled: !!branchId && !!id,
    });
};

export const useCreateSale = (): UseMutationResult<Sale, Error, { branchId: string; data: CreateSaleRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => saleApi.createSale(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['sales', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['inventory', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['shifts', variables.branchId] });
        },
    });
};

export const useSaleByReceipt = (branchId: string, receiptNumber: string): UseQueryResult<Sale, Error> => {
    return useQuery({
        queryKey: ['sales', branchId, 'receipt', receiptNumber],
        queryFn: () => saleApi.getSaleByReceipt(branchId, receiptNumber),
        enabled: !!branchId && !!receiptNumber,
    });
};

export const useRefundSale = (): UseMutationResult<Sale, Error, { branchId: string; id: string; reason?: string }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id, reason }) => saleApi.refundSale(branchId, id, reason),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['sales', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['inventory', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
        },
    });
};
