/**
 * Customer related types
 */

import { BaseModel } from './common';
import { Branch } from './branch';

export interface Customer extends BaseModel {
    branch_id: string;
    first_name: string;
    last_name: string;
    email?: string;
    phone: string;
    loyalty_points: number;
    total_spend: number;
    branch?: Branch;
}

export interface CreateCustomerRequest {
    first_name: string;
    last_name: string;
    email?: string;
    phone: string;
}

export interface UpdateCustomerRequest {
    first_name?: string;
    last_name?: string;
    email?: string;
    phone?: string;
}

export interface CustomerListParams {
    search?: string;
    min_loyalty_points?: number;
    min_total_spend?: number;
}
