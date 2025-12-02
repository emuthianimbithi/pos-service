/**
 * Notification and Alert React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { notificationApi } from '../api';
import {
    Alert,
    Notification,
    AlertListParams,
    NotificationListParams,
    MarkAlertReadRequest,
    MarkNotificationReadRequest,
} from '../types';

// Alerts
export const useAlerts = (params?: AlertListParams): UseQueryResult<Alert[], Error> => {
    return useQuery({
        queryKey: ['alerts', params],
        queryFn: () => notificationApi.getAlerts(params),
    });
};

export const useAlert = (id: string): UseQueryResult<Alert, Error> => {
    return useQuery({
        queryKey: ['alerts', id],
        queryFn: () => notificationApi.getAlert(id),
        enabled: !!id,
    });
};

export const useMarkAlertsRead = (): UseMutationResult<void, Error, MarkAlertReadRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: notificationApi.markAlertsRead,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['alerts'] });
            queryClient.invalidateQueries({ queryKey: ['unreadAlertCount'] });
        },
    });
};

export const useMarkAlertRead = (): UseMutationResult<Alert, Error, string> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: notificationApi.markAlertRead,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['alerts'] });
            queryClient.invalidateQueries({ queryKey: ['unreadAlertCount'] });
        },
    });
};

export const useUnreadAlertCount = (): UseQueryResult<{ count: number }, Error> => {
    return useQuery({
        queryKey: ['unreadAlertCount'],
        queryFn: notificationApi.getUnreadAlertCount,
        refetchInterval: 30000, // Refetch every 30 seconds
    });
};

// Notifications
export const useNotifications = (params?: NotificationListParams): UseQueryResult<Notification[], Error> => {
    return useQuery({
        queryKey: ['notifications', params],
        queryFn: () => notificationApi.getNotifications(params),
    });
};

export const useNotification = (id: string): UseQueryResult<Notification, Error> => {
    return useQuery({
        queryKey: ['notifications', id],
        queryFn: () => notificationApi.getNotification(id),
        enabled: !!id,
    });
};

export const useMarkNotificationsRead = (): UseMutationResult<void, Error, MarkNotificationReadRequest> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: notificationApi.markNotificationsRead,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['notifications'] });
            queryClient.invalidateQueries({ queryKey: ['unreadNotificationCount'] });
        },
    });
};

export const useMarkNotificationRead = (): UseMutationResult<Notification, Error, string> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: notificationApi.markNotificationRead,
        onSuccess: () => {
            queryClient.invalidateQueries({ queryKey: ['notifications'] });
            queryClient.invalidateQueries({ queryKey: ['unreadNotificationCount'] });
        },
    });
};

export const useUnreadNotificationCount = (): UseQueryResult<{ count: number }, Error> => {
    return useQuery({
        queryKey: ['unreadNotificationCount'],
        queryFn: notificationApi.getUnreadNotificationCount,
        refetchInterval: 30000, // Refetch every 30 seconds
    });
};
