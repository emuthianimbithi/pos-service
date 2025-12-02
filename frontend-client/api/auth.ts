/**
 * Authentication API functions
 */

import { apiClient, setToken, setUser, clearAuth } from './client';
import {
    LoginRequest,
    LoginResponse,
    RegisterRequest,
    RegisterResponse,
    ChangePasswordRequest,
    ResetPasswordRequest,
    ConfirmResetPasswordRequest,
    AuthUser,
} from '../types';

export const authApi = {
    /**
     * Login with email and password
     */
    login: async (data: LoginRequest): Promise<LoginResponse> => {
        const response = await apiClient.post<{ success: boolean; data: LoginResponse }>('/auth/login', data);
        const loginData = response.data.data!;
        setToken(loginData.token);
        setUser(loginData.user);
        return loginData;
    },

    /**
     * Register a new business and admin user
     */
    register: async (data: RegisterRequest): Promise<RegisterResponse> => {
        const response = await apiClient.post<RegisterResponse>('/auth/register', data);
        setToken(response.data.token);
        setUser(response.data.user);
        return response.data;
    },

    /**
     * Logout the current user
     */
    logout: async (): Promise<void> => {
        try {
            await apiClient.post('/auth/logout');
        } finally {
            clearAuth();
        }
    },

    /**
     * Get current authenticated user
     */
    getCurrentUser: async (): Promise<AuthUser> => {
        const response = await apiClient.get<AuthUser>('/auth/me');
        setUser(response.data);
        return response.data;
    },

    /**
     * Change password for current user
     */
    changePassword: async (data: ChangePasswordRequest): Promise<void> => {
        await apiClient.post('/auth/change-password', data);
    },

    /**
     * Request password reset
     */
    requestPasswordReset: async (data: ResetPasswordRequest): Promise<void> => {
        await apiClient.post('/auth/reset-password', data);
    },

    /**
     * Confirm password reset with token
     */
    confirmPasswordReset: async (data: ConfirmResetPasswordRequest): Promise<void> => {
        await apiClient.post('/auth/confirm-reset-password', data);
    },
};
