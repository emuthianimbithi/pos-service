/**
 * Report React Query hooks
 */

import { useQuery, UseQueryResult } from '@tanstack/react-query';
import { reportApi } from '../api';
import {
    SalesReport,
    InventoryReport,
    RevenueReport,
    CustomerReport,
    ReportParams,
} from '../types';

export const useSalesReport = (branchId: string, params: ReportParams): UseQueryResult<SalesReport, Error> => {
    return useQuery({
        queryKey: ['reports', 'sales', branchId, params],
        queryFn: () => reportApi.getSalesReport(branchId, params),
        enabled: !!branchId && !!params.start_date && !!params.end_date,
    });
};

export const useInventoryReport = (branchId: string): UseQueryResult<InventoryReport, Error> => {
    return useQuery({
        queryKey: ['reports', 'inventory', branchId],
        queryFn: () => reportApi.getInventoryReport(branchId),
        enabled: !!branchId,
    });
};

export const useRevenueReport = (branchId: string, params: ReportParams): UseQueryResult<RevenueReport, Error> => {
    return useQuery({
        queryKey: ['reports', 'revenue', branchId, params],
        queryFn: () => reportApi.getRevenueReport(branchId, params),
        enabled: !!branchId && !!params.start_date && !!params.end_date,
    });
};

export const useCustomerReport = (branchId: string, params?: Partial<ReportParams>): UseQueryResult<CustomerReport, Error> => {
    return useQuery({
        queryKey: ['reports', 'customers', branchId, params],
        queryFn: () => reportApi.getCustomerReport(branchId, params),
        enabled: !!branchId,
    });
};

// Export helpers
export const exportSalesReportCSV = async (branchId: string, params: ReportParams): Promise<void> => {
    const blob = await reportApi.exportSalesReportCSV(branchId, params);
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `sales-report-${params.start_date}-${params.end_date}.csv`;
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
};

export const exportInventoryReportCSV = async (branchId: string): Promise<void> => {
    const blob = await reportApi.exportInventoryReportCSV(branchId);
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `inventory-report-${new Date().toISOString().split('T')[0]}.csv`;
    document.body.appendChild(a);
    a.click();
    window.URL.revokeObjectURL(url);
    document.body.removeChild(a);
};
