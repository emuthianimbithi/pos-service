/**
 * Authentication React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { authApi } from '../api';
import {
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse,
    AuthUser,
    ChangePasswordRequest,
    ResetPasswordRequest,
    ConfirmResetPasswordRequest,
} from '../types';

/**
 * Get current authenticated user
 */
export const useCurrentUser = (): UseQueryResult<AuthUser, Error> => {
    return useQuery({
        queryKey: ['auth', 'currentUser'],
        queryFn: authApi.getCurrentUser,
        retry: false,
    });
};

/**
 * Login mutation
 */
export const useLogin = (): UseMutationResult<LoginResponse, Error, LoginRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: authApi.login,
        onSuccess: (data) => {
            queryClient.setQueryData(['auth', 'currentUser'], data.user);
        },
    });
};

/**
 * Register mutation
 */
export const useRegister = (): UseMutationResult<RegisterResponse, Error, RegisterRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: authApi.register,
        onSuccess: (data) => {
            queryClient.setQueryData(['auth', 'currentUser'], data.user);
        },
    });
};

/**
 * Logout mutation
 */
export const useLogout = (): UseMutationResult<void, Error, void> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: authApi.logout,
        onSuccess: () => {
            queryClient.clear();
        },
    });
};

/**
 * Change password mutation
 */
export const useChangePassword = (): UseMutationResult<void, Error, ChangePasswordRequest> => {
    return useMutation({
        mutationFn: authApi.changePassword,
    });
};

/**
 * Request password reset mutation
 */
export const useRequestPasswordReset = (): UseMutationResult<void, Error, ResetPasswordRequest> => {
    return useMutation({
        mutationFn: authApi.requestPasswordReset,
    });
};

/**
 * Confirm password reset mutation
 */
export const useConfirmPasswordReset = (): UseMutationResult<void, Error, ConfirmResetPasswordRequest> => {
    return useMutation({
        mutationFn: authApi.confirmPasswordReset,
    });
};
