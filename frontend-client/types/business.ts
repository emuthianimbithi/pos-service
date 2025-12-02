/**
 * Business related types
 */

import { BaseModel } from './common';
import { Branch } from './branch';
import { User } from './user';

export interface Business extends BaseModel {
    name: string;
    email: string;
    phone?: string;
    address?: string;
    tax_id?: string;
    logo_url?: string;
    is_active: boolean;
    branches?: Branch[];
    users?: User[];
}

export interface CreateBusinessRequest {
    name: string;
    email: string;
    phone?: string;
    address?: string;
    tax_id?: string;
}

export interface UpdateBusinessRequest {
    name?: string;
    email?: string;
    phone?: string;
    address?: string;
    tax_id?: string;
    logo_url?: string;
    is_active?: boolean;
}
