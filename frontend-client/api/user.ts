/**
 * User API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    User,
    CreateUserRequest,
    UpdateUserRequest,
    UserListParams,
    PaginatedResponse,
} from '../types';

export const userApi = {
    /**
     * Get all users with optional filters
     */
    getUsers: async (params?: UserListParams): Promise<User[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<User[]>(`/admin/users${query}`);
        return response.data;
    },

    /**
     * Get user by ID
     */
    getUser: async (id: string): Promise<User> => {
        const response = await apiClient.get<User>(`/admin/users/${id}`);
        return response.data;
    },

    /**
     * Create a new user
     */
    createUser: async (data: CreateUserRequest): Promise<User> => {
        const response = await apiClient.post<User>('/admin/users', data);
        return response.data;
    },

    /**
     * Update user by ID
     */
    updateUser: async (id: string, data: UpdateUserRequest): Promise<User> => {
        const response = await apiClient.put<User>(`/admin/users/${id}`, data);
        return response.data;
    },

    /**
     * Delete user by ID
     */
    deleteUser: async (id: string): Promise<void> => {
        await apiClient.delete(`/admin/users/${id}`);
    },

    /**
     * Activate user
     */
    activateUser: async (id: string): Promise<User> => {
        const response = await apiClient.post<User>(`/admin/users/${id}/activate`);
        return response.data;
    },

    /**
     * Deactivate user
     */
    deactivateUser: async (id: string): Promise<User> => {
        const response = await apiClient.post<User>(`/admin/users/${id}/deactivate`);
        return response.data;
    },
};
