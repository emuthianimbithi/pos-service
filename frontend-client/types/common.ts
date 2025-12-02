/**
 * Common TypeScript types for POS Service
 */

export interface BaseModel {
  id: string; // UUID
  created_at: string; // ISO 8601 timestamp
  updated_at: string; // ISO 8601 timestamp
  deleted_at?: string | null; // ISO 8601 timestamp
}

export interface TenantModel {
  business_id: string; // UUID
}

export interface PaginationParams {
  page?: number;
  page_size?: number;
  sort_by?: string;
  sort_order?: 'asc' | 'desc';
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  page_size: number;
  total_pages: number;
}

// Backend response wrapper (matches pkg/response/response.go)
export interface BackendResponse<T = any> {
  success: boolean;
  message?: string;
  data?: T;
  error?: string;
  meta?: PaginationMeta;
}

export interface PaginationMeta {
  current_page: number;
  per_page: number;
  total: number;
  total_pages: number;
}

export interface ApiResponse<T> {
  success: boolean;
  data?: T;
  message?: string;
  error?: string;
}

export interface ApiError {
  error: string;
  message?: string;
  status?: number;
  details?: Record<string, string>; // For validation errors
}

// User roles
export type UserRole = 'admin' | 'manager' | 'cashier' | 'stock_keeper';

// Payment methods
export type PaymentMethod = 'cash' | 'card' | 'mpesa';

// Transaction types
export type TransactionType = 'adjustment' | 'sale' | 'return' | 'transfer_in' | 'transfer_out';

// Alert priority
export type AlertPriority = 'low' | 'medium' | 'high' | 'critical';

// Notification types
export type NotificationType = 'success' | 'error' | 'info' | 'warning';

// Alert types
export type AlertType = 'stock_low' | 'security' | 'system';

// Stock status
export type StockStatus = 'in_stock' | 'low_stock' | 'out_of_stock';

// Shift status
export type ShiftStatus = 'open' | 'closed';

// Sale status
export type SaleStatus = 'completed' | 'pending' | 'cancelled' | 'refunded';
