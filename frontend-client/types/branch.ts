/**
 * Branch related types
 */

import { BaseModel } from './common';
import { Business } from './business';

export interface Branch extends BaseModel {
    business_id: string;
    name: string;
    code: string;
    phone?: string;
    address?: string;
    is_active: boolean;
    business?: Business;
}

export interface CreateBranchRequest {
    name: string;
    code: string;
    phone?: string;
    address?: string;
}

export interface UpdateBranchRequest {
    name?: string;
    code?: string;
    phone?: string;
    address?: string;
    is_active?: boolean;
}

export interface BranchListParams {
    is_active?: boolean;
    search?: string;
}
