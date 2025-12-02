/**
 * Base API Client Configuration
 * Axios instance with authentication and error handling
 */

import axios, { AxiosInstance, AxiosError, InternalAxiosRequestConfig } from 'axios';
import { ApiError } from '../types';

// Default API base URL (can be overridden via environment variable)
const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api';

// Token storage keys
const TOKEN_KEY = 'pos_auth_token';
const USER_KEY = 'pos_auth_user';

/**
 * Create axios instance with default configuration
 */
export const apiClient: AxiosInstance = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
    timeout: 30000, // 30 seconds
});

/**
 * Request interceptor to add auth token
 */
apiClient.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        const token = getToken();
        if (token && config.headers) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

/**
 * Response interceptor to unwrap backend response and handle errors
 */
apiClient.interceptors.response.use(
    (response) => {
        // Unwrap the backend response structure
        // Backend sends: { success: true, data: {...}, message: "..." }
        // We want to return just the data for easier consumption
        if (response.data && typeof response.data === 'object' && 'success' in response.data) {
            // Keep the full response structure but make data easily accessible
            return {
                ...response,
                // Original wrapped response available as response.data
                // Unwrapped data available as response.data.data
            };
        }
        return response;
    },
    (error: AxiosError<ApiError>) => {
        // Handle 401 Unauthorized - clear auth and redirect to login
        if (error.response?.status === 401) {
            clearAuth();
            window.location.href = '/login';
        }

        // Format error response from backend
        const responseData: any = error.response?.data;
        const apiError: ApiError = {
            error: responseData?.error || error.message || 'An error occurred',
            message: responseData?.message,
            status: error.response?.status,
            details: responseData?.details, // Validation error details
        };

        return Promise.reject(apiError);
    }
);

/**
 * Authentication helpers
 */

export const setToken = (token: string): void => {
    localStorage.setItem(TOKEN_KEY, token);
};

export const getToken = (): string | null => {
    return localStorage.getItem(TOKEN_KEY);
};

export const setUser = (user: any): void => {
    localStorage.setItem(USER_KEY, JSON.stringify(user));
};

export const getUser = (): any | null => {
    const user = localStorage.getItem(USER_KEY);
    return user ? JSON.parse(user) : null;
};

export const clearAuth = (): void => {
    localStorage.removeItem(TOKEN_KEY);
    localStorage.removeItem(USER_KEY);
};

export const isAuthenticated = (): boolean => {
    return !!getToken();
};

/**
 * Helper function to build query string from params
 */
export const buildQueryString = (params?: Record<string, any>): string => {
    if (!params) return '';

    const filtered = Object.entries(params)
        .filter(([_, value]) => value !== undefined && value !== null && value !== '')
        .map(([key, value]) => `${encodeURIComponent(key)}=${encodeURIComponent(value)}`)
        .join('&');

    return filtered ? `?${filtered}` : '';
};
