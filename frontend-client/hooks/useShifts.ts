/**
 * Shift React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { shiftApi } from '../api';
import {
    Shift,
    CreateShiftRequest,
    CloseShiftRequest,
    ShiftListParams,
} from '../types';

export const useShifts = (branchId: string, params?: ShiftListParams): UseQueryResult<Shift[], Error> => {
    return useQuery({
        queryKey: ['shifts', branchId, params],
        queryFn: () => shiftApi.getShifts(branchId, params),
        enabled: !!branchId,
    });
};

export const useShift = (branchId: string, id: string): UseQueryResult<Shift, Error> => {
    return useQuery({
        queryKey: ['shifts', branchId, id],
        queryFn: () => shiftApi.getShift(branchId, id),
        enabled: !!branchId && !!id,
    });
};

export const useStartShift = (): UseMutationResult<Shift, Error, { branchId: string; data: CreateShiftRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => shiftApi.startShift(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['shifts', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['activeShift', variables.branchId] });
        },
    });
};

export const useCloseShift = (): UseMutationResult<Shift, Error, { branchId: string; id: string; data: CloseShiftRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id, data }) => shiftApi.closeShift(branchId, id, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['shifts', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['shifts', variables.branchId, variables.id] });
            queryClient.invalidateQueries({ queryKey: ['activeShift', variables.branchId] });
        },
    });
};

export const useActiveShift = (branchId: string): UseQueryResult<Shift | null, Error> => {
    return useQuery({
        queryKey: ['activeShift', branchId],
        queryFn: () => shiftApi.getActiveShift(branchId),
        enabled: !!branchId,
    });
};

export const useShiftSummary = (branchId: string, id: string): UseQueryResult<Shift, Error> => {
    return useQuery({
        queryKey: ['shifts', branchId, id, 'summary'],
        queryFn: () => shiftApi.getShiftSummary(branchId, id),
        enabled: !!branchId && !!id,
    });
};
