/**
 * Business React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { businessApi } from '../api';
import {
    Business,
    UpdateBusinessRequest,
} from '../types';

/**
 * Get current business
 */
export const useBusiness = (): UseQueryResult<Business, Error> => {
    return useQuery({
        queryKey: ['business'],
        queryFn: businessApi.getBusiness,
    });
};

/**
 * Update business mutation
 */
export const useUpdateBusiness = (): UseMutationResult<Business, Error, UpdateBusinessRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: businessApi.updateBusiness,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['business'] });
        },
    });
};

/**
 * Upload business logo mutation
 */
export const useUploadBusinessLogo = (): UseMutationResult<{ logo_url: string }, Error, File> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: businessApi.uploadLogo,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['business'] });
        },
    });
};
