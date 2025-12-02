/**
 * Authentication related types
 */

import { UserRole } from './common';

export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    token: string;
    user: AuthUser;
}

export interface RegisterRequest {
    email: string;
    password: string;
    first_name: string;
    last_name: string;
    phone?: string;
    business_name: string;
    business_email: string;
    business_phone?: string;
}

export interface RegisterResponse {
    token: string;
    user: AuthUser;
}

export interface AuthUser {
    id: string;
    email: string;
    first_name: string;
    last_name: string;
    phone?: string;
    role: UserRole;
    business_id: string;
    branch_id?: string;
    is_active: boolean;
}

export interface ChangePasswordRequest {
    old_password: string;
    new_password: string;
}

export interface ResetPasswordRequest {
    email: string;
}

export interface ConfirmResetPasswordRequest {
    token: string;
    new_password: string;
}
