/**
 * Shift related types
 */

import { BaseModel, ShiftStatus } from './common';
import { Branch } from './branch';
import { User } from './user';

export interface Shift extends BaseModel {
    branch_id: string;
    user_id: string;
    start_time: string;
    end_time?: string;
    opening_balance: number;
    closing_balance?: number;
    cash_sales: number;
    mpesa_sales: number;
    card_sales: number;
    total_sales: number;
    cash_variance: number;
    notes?: string;
    status: ShiftStatus;
    branch?: Branch;
    user?: User;
}

export interface CreateShiftRequest {
    opening_balance: number;
    notes?: string;
}

export interface CloseShiftRequest {
    closing_balance: number;
    notes?: string;
}

export interface ShiftListParams {
    user_id?: string;
    status?: ShiftStatus;
    start_date?: string;
    end_date?: string;
}
