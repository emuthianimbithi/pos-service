/**
 * Notification and Alert related types
 */

import { BaseModel, AlertPriority, NotificationType, AlertType } from './common';

export interface Alert extends BaseModel {
    business_id: string;
    branch_id?: string;
    title: string;
    message: string;
    type: AlertType;
    priority: AlertPriority;
    action_link?: string;
    target_user_id?: string;
    target_role?: string;
    is_read: boolean;
    read_at?: string;
    expires_at?: string;
}

export interface Notification extends BaseModel {
    business_id: string;
    user_id: string;
    title: string;
    message: string;
    type: NotificationType;
    is_read: boolean;
}

export interface MarkAlertReadRequest {
    alert_ids: string[];
}

export interface MarkNotificationReadRequest {
    notification_ids: string[];
}

export interface AlertListParams {
    is_read?: boolean;
    type?: AlertType;
    priority?: AlertPriority;
    target_role?: string;
}

export interface NotificationListParams {
    is_read?: boolean;
    type?: NotificationType;
}
