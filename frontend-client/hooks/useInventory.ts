/**
 * Inventory React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { inventoryApi } from '../api';
import {
    InventoryTransaction,
    CreateInventoryTransactionRequest,
    InventoryListParams,
    StockAdjustmentRequest,
    LowStockAlert,
} from '../types';

export const useInventoryTransactions = (branchId: string, params?: InventoryListParams): UseQueryResult<InventoryTransaction[], Error> => {
    return useQuery({
        queryKey: ['inventory', branchId, params],
        queryFn: () => inventoryApi.getTransactions(branchId, params),
        enabled: !!branchId,
    });
};

export const useInventoryTransaction = (branchId: string, id: string): UseQueryResult<InventoryTransaction, Error> => {
    return useQuery({
        queryKey: ['inventory', branchId, id],
        queryFn: () => inventoryApi.getTransaction(branchId, id),
        enabled: !!branchId && !!id,
    });
};

export const useCreateInventoryTransaction = (): UseMutationResult<InventoryTransaction, Error, { branchId: string; data: CreateInventoryTransactionRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => inventoryApi.createTransaction(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['inventory', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
        },
    });
};

export const useAdjustStock = (): UseMutationResult<InventoryTransaction, Error, { branchId: string; data: StockAdjustmentRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => inventoryApi.adjustStock(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['inventory', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['lowStockAlerts', variables.branchId] });
        },
    });
};

export const useLowStockAlerts = (branchId: string): UseQueryResult<LowStockAlert[], Error> => {
    return useQuery({
        queryKey: ['lowStockAlerts', branchId],
        queryFn: () => inventoryApi.getLowStockAlerts(branchId),
        enabled: !!branchId,
    });
};

export const useProductInventory = (branchId: string, productId: string): UseQueryResult<InventoryTransaction[], Error> => {
    return useQuery({
        queryKey: ['inventory', branchId, 'product', productId],
        queryFn: () => inventoryApi.getProductInventory(branchId, productId),
        enabled: !!branchId && !!productId,
    });
};
