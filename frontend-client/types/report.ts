/**
 * Report related types
 */

export interface SalesReport {
    start_date: string;
    end_date: string;
    total_sales: number;
    total_orders: number;
    cash_sales: number;
    card_sales: number;
    mpesa_sales: number;
    average_order: number;
    top_products: ProductSalesReport[];
    sales_by_day: DailySalesReport[];
}

export interface ProductSalesReport {
    product_id: string;
    product_name: string;
    quantity: number;
    revenue: number;
}

export interface DailySalesReport {
    date: string;
    sales: number;
    orders: number;
}

export interface InventoryReport {
    total_products: number;
    low_stock_products: number;
    out_of_stock: number;
    total_value: number;
    products: ProductInventoryReport[];
}

export interface ProductInventoryReport {
    product_id: string;
    product_name: string;
    variant_name: string;
    sku: string;
    stock: number;
    low_stock_threshold: number;
    status: string;
    value: number;
}

export interface RevenueReport {
    period: string;
    total_revenue: number;
    total_cost: number;
    gross_profit: number;
    profit_margin: number;
    daily_breakdown: DailyRevenueReport[];
}

export interface DailyRevenueReport {
    date: string;
    revenue: number;
    cost: number;
    profit: number;
    profit_margin: number;
}

export interface CustomerReport {
    total_customers: number;
    top_customers: CustomerSalesReport[];
}

export interface CustomerSalesReport {
    customer_id: string;
    customer_name: string;
    total_purchases: number;
    order_count: number;
    loyalty_points: number;
}

export interface ReportParams {
    start_date: string;
    end_date: string;
    branch_id?: string;
}
