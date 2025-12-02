/**
 * Customer React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { customerApi } from '../api';
import {
    Customer,
    CreateCustomerRequest,
    UpdateCustomerRequest,
    CustomerListParams,
} from '../types';

export const useCustomers = (branchId: string, params?: CustomerListParams): UseQueryResult<Customer[], Error> => {
    return useQuery({
        queryKey: ['customers', branchId, params],
        queryFn: () => customerApi.getCustomers(branchId, params),
        enabled: !!branchId,
    });
};

export const useCustomer = (branchId: string, id: string): UseQueryResult<Customer, Error> => {
    return useQuery({
        queryKey: ['customers', branchId, id],
        queryFn: () => customerApi.getCustomer(branchId, id),
        enabled: !!branchId && !!id,
    });
};

export const useCreateCustomer = (): UseMutationResult<Customer, Error, { branchId: string; data: CreateCustomerRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => customerApi.createCustomer(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['customers', variables.branchId] });
        },
    });
};

export const useUpdateCustomer = (): UseMutationResult<Customer, Error, { branchId: string; id: string; data: UpdateCustomerRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id, data }) => customerApi.updateCustomer(branchId, id, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['customers', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['customers', variables.branchId, variables.id] });
        },
    });
};

export const useDeleteCustomer = (): UseMutationResult<void, Error, { branchId: string; id: string }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id }) => customerApi.deleteCustomer(branchId, id),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['customers', variables.branchId] });
        },
    });
};

export const useSearchCustomerByPhone = (branchId: string, phone: string): UseQueryResult<Customer[], Error> => {
    return useQuery({
        queryKey: ['customers', branchId, 'search', phone],
        queryFn: () => customerApi.searchByPhone(branchId, phone),
        enabled: !!branchId && !!phone && phone.length >= 3,
    });
};
