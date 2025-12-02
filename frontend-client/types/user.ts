/**
 * User related types
 */

import { BaseModel, UserRole } from './common';
import { Business } from './business';
import { Branch } from './branch';

export interface User extends BaseModel {
    business_id: string;
    branch_id?: string;
    email: string;
    first_name: string;
    last_name: string;
    phone?: string;
    role: UserRole;
    is_active: boolean;
    business?: Business;
    branch?: Branch;
}

export interface CreateUserRequest {
    email: string;
    password: string;
    first_name: string;
    last_name: string;
    phone?: string;
    role: UserRole;
    branch_id?: string;
}

export interface UpdateUserRequest {
    first_name?: string;
    last_name?: string;
    phone?: string;
    role?: UserRole;
    branch_id?: string;
    is_active?: boolean;
}

export interface UserListParams {
    role?: UserRole;
    branch_id?: string;
    is_active?: boolean;
    search?: string;
}
