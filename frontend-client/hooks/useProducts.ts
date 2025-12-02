/**
 * Product React Query hooks
 */

import { useMutation, useQuery, useQueryClient, UseMutationResult, UseQueryResult } from '@tanstack/react-query';
import { productApi } from '../api';
import {
    Product,
    ProductVariant,
    Category,
    CreateProductRequest,
    UpdateProductRequest,
    CreateProductVariantRequest,
    UpdateProductVariantRequest,
    CreateCategoryRequest,
    UpdateCategoryRequest,
    ProductListParams,
} from '../types';

// Categories
export const useCategories = (branchId: string): UseQueryResult<Category[], Error> => {
    return useQuery({
        queryKey: ['categories', branchId],
        queryFn: () => productApi.getCategories(branchId),
        enabled: !!branchId,
    });
};

export const useCategory = (branchId: string, id: string): UseQueryResult<Category, Error> => {
    return useQuery({
        queryKey: ['categories', branchId, id],
        queryFn: () => productApi.getCategory(branchId, id),
        enabled: !!branchId && !!id,
    });
};

export const useCreateCategory = (): UseMutationResult<Category, Error, { branchId: string; data: CreateCategoryRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => productApi.createCategory(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['categories', variables.branchId] });
        },
    });
};

export const useUpdateCategory = (): UseMutationResult<Category, Error, { branchId: string; id: string; data: UpdateCategoryRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id, data }) => productApi.updateCategory(branchId, id, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['categories', variables.branchId] });
        },
    });
};

export const useDeleteCategory = (): UseMutationResult<void, Error, { branchId: string; id: string }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id }) => productApi.deleteCategory(branchId, id),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['categories', variables.branchId] });
        },
    });
};

// Products
export const useProducts = (branchId: string, params?: ProductListParams): UseQueryResult<Product[], Error> => {
    return useQuery({
        queryKey: ['products', branchId, params],
        queryFn: () => productApi.getProducts(branchId, params),
        enabled: !!branchId,
    });
};

export const useProduct = (branchId: string, id: string): UseQueryResult<Product, Error> => {
    return useQuery({
        queryKey: ['products', branchId, id],
        queryFn: () => productApi.getProduct(branchId, id),
        enabled: !!branchId && !!id,
    });
};

export const useCreateProduct = (): UseMutationResult<Product, Error, { branchId: string; data: CreateProductRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, data }) => productApi.createProduct(branchId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['inventory', variables.branchId] });
        },
    });
};

export const useUpdateProduct = (): UseMutationResult<Product, Error, { branchId: string; id: string; data: UpdateProductRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id, data }) => productApi.updateProduct(branchId, id, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId, variables.id] });
        },
    });
};

export const useDeleteProduct = (): UseMutationResult<void, Error, { branchId: string; id: string }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, id }) => productApi.deleteProduct(branchId, id),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId] });
        },
    });
};

// Product Variants
export const useProductVariants = (branchId: string, productId: string): UseQueryResult<ProductVariant[], Error> => {
    return useQuery({
        queryKey: ['productVariants', branchId, productId],
        queryFn: () => productApi.getProductVariants(branchId, productId),
        enabled: !!branchId && !!productId,
    });
};

export const useCreateProductVariant = (): UseMutationResult<ProductVariant, Error, { branchId: string; productId: string; data: CreateProductVariantRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, productId, data }) => productApi.createProductVariant(branchId, productId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['productVariants', variables.branchId, variables.productId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId, variables.productId] });
        },
    });
};

export const useUpdateProductVariant = (): UseMutationResult<ProductVariant, Error, { branchId: string; productId: string; variantId: string; data: UpdateProductVariantRequest }> => {
    const queryClient = useQueryClient();

    return useMutation({
        mutationFn: ({ branchId, productId, variantId, data }) => productApi.updateProductVariant(branchId, productId, variantId, data),
        onSuccess: (_, variables) => {
            queryClient.invalidateQueries({ queryKey: ['productVariants', variables.branchId, variables.productId] });
            queryClient.invalidateQueries({ queryKey: ['products', variables.branchId, variables.productId] });
        },
    });
};
