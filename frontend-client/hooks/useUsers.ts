/**
 * User Management React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { userApi } from '../api';
import {
    User,
    CreateUserRequest,
    UpdateUserRequest,
    UserListParams,
} from '../types';

/**
 * Get all users
 */
export const useUsers = (params?: UserListParams): UseQueryResult<User[], Error> => {
    return useQuery({
        queryKey: ['users', params],
        queryFn: () => userApi.getUsers(params),
    });
};

/**
 * Get user by ID
 */
export const useUser = (id: string): UseQueryResult<User, Error> => {
    return useQuery({
        queryKey: ['users', id],
        queryFn: () => userApi.getUser(id),
        enabled: !!id,
    });
};

/**
 * Create user mutation
 */
export const useCreateUser = (): UseMutationResult<User, Error, CreateUserRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: userApi.createUser,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
        },
    });
};

/**
 * Update user mutation
 */
export const useUpdateUser = (): UseMutationResult<User, Error, { id: string; data: UpdateUserRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ id, data }) => userApi.updateUser(id, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
            queryClient.invalidateQueries({ queryKey: ['users', variables.id] });
        },
    });
};

/**
 * Delete user mutation
 */
export const useDeleteUser = (): UseMutationResult<void, Error, string> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: userApi.deleteUser,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
        },
    });
};

/**
 * Activate user mutation
 */
export const useActivateUser = (): UseMutationResult<User, Error, string> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: userApi.activateUser,
        onSuccess: (_, id) => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
            queryClient.invalidateQueries({ queryKey: ['users', id] });
        },
    });
};

/**
 * Deactivate user mutation
 */
export const useDeactivateUser = (): UseMutationResult<User, Error, string> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: userApi.deactivateUser,
        onSuccess: (_, id) => {
            queryClient.invalidateQueries({ queryKey: ['users'] });
            queryClient.invalidateQueries({ queryKey: ['users', id] });
        },
    });
};
