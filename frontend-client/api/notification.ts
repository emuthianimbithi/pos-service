/**
 * Notification and Alert API functions
 */

import { apiClient, buildQueryString } from './client';
import {
    Alert,
    Notification,
    AlertListParams,
    NotificationListParams,
    MarkAlertReadRequest,
    MarkNotificationReadRequest,
} from '../types';

export const notificationApi = {
    // Alerts
    /**
     * Get all alerts with optional filters
     */
    getAlerts: async (params?: AlertListParams): Promise<Alert[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Alert[]>(`/alerts${query}`);
        return response.data;
    },

    /**
     * Get alert by ID
     */
    getAlert: async (id: string): Promise<Alert> => {
        const response = await apiClient.get<Alert>(`/alerts/${id}`);
        return response.data;
    },

    /**
     * Mark alerts as read
     */
    markAlertsRead: async (data: MarkAlertReadRequest): Promise<void> => {
        await apiClient.post('/alerts/mark-read', data);
    },

    /**
     * Mark single alert as read
     */
    markAlertRead: async (id: string): Promise<Alert> => {
        const response = await apiClient.post<Alert>(`/alerts/${id}/read`);
        return response.data;
    },

    /**
     * Get unread alert count
     */
    getUnreadAlertCount: async (): Promise<{ count: number }> => {
        const response = await apiClient.get<{ count: number }>('/alerts/unread/count');
        return response.data;
    },

    // Notifications
    /**
     * Get all notifications with optional filters
     */
    getNotifications: async (params?: NotificationListParams): Promise<Notification[]> => {
        const query = buildQueryString(params);
        const response = await apiClient.get<Notification[]>(`/notifications${query}`);
        return response.data;
    },

    /**
     * Get notification by ID
     */
    getNotification: async (id: string): Promise<Notification> => {
        const response = await apiClient.get<Notification>(`/notifications/${id}`);
        return response.data;
    },

    /**
     * Mark notifications as read
     */
    markNotificationsRead: async (data: MarkNotificationReadRequest): Promise<void> => {
        await apiClient.post('/notifications/mark-read', data);
    },

    /**
     * Mark single notification as read
     */
    markNotificationRead: async (id: string): Promise<Notification> => {
        const response = await apiClient.post<Notification>(`/notifications/${id}/read`);
        return response.data;
    },

    /**
     * Get unread notification count
     */
    getUnreadNotificationCount: async (): Promise<{ count: number }> => {
        const response = await apiClient.get<{ count: number }>('/notifications/unread/count');
        return response.data;
    },
};
