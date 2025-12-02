/**
 * Branch React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { branchApi } from '../api';
import {
    Branch,
    CreateBranchRequest,
    UpdateBranchRequest,
    BranchListParams,
} from '../types';

/**
 * Get all branches
 */
export const useBranches = (params?: BranchListParams): UseQueryResult<Branch[], Error> => {
    return useQuery({
        queryKey: ['branches', params],
        queryFn: () => branchApi.getBranches(params),
    });
};

/**
 * Get branch by ID
 */
export const useBranch = (id: string): UseQueryResult<Branch, Error> => {
    return useQuery({
        queryKey: ['branches', id],
        queryFn: () => branchApi.getBranch(id),
        enabled: !!id,
    });
};

/**
 * Create branch mutation
 */
export const useCreateBranch = (): UseMutationResult<Branch, Error, CreateBranchRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: branchApi.createBranch,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['branches'] });
        },
    });
};

/**
 * Update branch mutation
 */
export const useUpdateBranch = (): UseMutationResult<Branch, Error, { id: string; data: UpdateBranchRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, data }) => branchApi.updateBranch(id, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['branches'] });
            queryClient.invalidateQueries({ queryKey: ['branches', variables.id] });
        },
    });
};

/**
 * Delete branch mutation
 */
export const useDeleteBranch = (): UseMutationResult<void, Error, string> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: branchApi.deleteBranch,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['branches'] });
        },
    });
};
