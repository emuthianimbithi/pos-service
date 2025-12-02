# POS Service TypeScript Client

A comprehensive TypeScript client library for the POS Service API, featuring type-safe API functions and React Query hooks for seamless integration with Vite/React applications.

## Features

✨ **Full TypeScript Support** - Complete type definitions for all API models
🔐 **Authentication Handling** - Automatic token management and refresh
📦 **React Query Integration** - Pre-built hooks with caching and state management
🎯 **Type-Safe API Calls** - Axios-based client with full IntelliSense support
⚡ **Optimized Performance** - Automatic cache invalidation and refetching

## Installation

### In Your Vite Project

```bash
# Install dependencies
npm install @tanstack/react-query axios

# Copy the frontend-client folder to your project
cp -r frontend-client /path/to/your/vite-project/src/
```

## Quick Start

### 1. Setup React Query

In your `main.tsx` or `App.tsx`:

```typescript
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      retry: 1,
      refetchOnWindowFocus: false,
    },
  },
});

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      {/* Your app components */}
    </QueryClientProvider>
  );
}
```

### 2. Configure API Base URL

Create a `.env` file in your project root:

```env
VITE_API_BASE_URL=http://localhost:8080/api
```

### 3. Use the Hooks in Your Components

#### Authentication

```typescript
import { useLogin, useCurrentUser, useLogout } from '@/frontend-client';

function LoginPage() {
  const login = useLogin();
  const { data: user } = useCurrentUser();
  const logout = useLogout();

  const handleLogin = async () => {
    try {
      await login.mutateAsync({
        email: 'user@example.com',
        password: 'password123',
      });
    } catch (error) {
      console.error('Login failed:', error);
    }
  };

  return (
    <div>
      {user ? (
        <>
          <p>Welcome, {user.first_name}!</p>
          <button onClick={() => logout.mutate()}>Logout</button>
        </>
      ) : (
        <button onClick={handleLogin}>Login</button>
      )}
    </div>
  );
}
```

#### Products Management

```typescript
import { useProducts, useCreateProduct, useUpdateProduct } from '@/frontend-client';

function ProductsPage() {
  const branchId = 'your-branch-id';
  const { data: products, isLoading } = useProducts(branchId);
  const createProduct = useCreateProduct();

  const handleCreateProduct = async () => {
    await createProduct.mutateAsync({
      branchId,
      data: {
        name: 'New Product',
        description: 'Product description',
        track_inventory: true,
        variants: [
          {
            name: 'Default',
            sku: 'SKU-001',
            price: 99.99,
            cost: 50.00,
            stock: 100,
          },
        ],
      },
    });
  };

  if (isLoading) return <div>Loading...</div>;

  return (
    <div>
      <button onClick={handleCreateProduct}>Add Product</button>
      {products?.map((product) => (
        <div key={product.id}>{product.name}</div>
      ))}
    </div>
  );
}
```

#### Sales Processing

```typescript
import { useCreateSale, useActiveShift } from '@/frontend-client';

function POSInterface() {
  const branchId = 'your-branch-id';
  const { data: activeShift } = useActiveShift(branchId);
  const createSale = useCreateSale();

  const handleCheckout = async (cartItems: any[]) => {
    await createSale.mutateAsync({
      branchId,
      data: {
        shift_id: activeShift?.id,
        payment_method: 'cash',
        items: cartItems.map((item) => ({
          product_id: item.product_id,
          variant_id: item.variant_id,
          quantity: item.quantity,
          unit_price: item.price,
        })),
      },
    });
  };

  return <div>{/* Your POS UI */}</div>;
}
```

#### Reports

```typescript
import { useSalesReport, exportSalesReportCSV } from '@/frontend-client';

function ReportsPage() {
  const branchId = 'your-branch-id';
  const { data: report } = useSalesReport(branchId, {
    start_date: '2024-01-01',
    end_date: '2024-01-31',
  });

  const handleExport = () => {
    exportSalesReportCSV(branchId, {
      start_date: '2024-01-01',
      end_date: '2024-01-31',
    });
  };

  return (
    <div>
      <h2>Total Sales: ${report?.total_sales}</h2>
      <button onClick={handleExport}>Export CSV</button>
    </div>
  );
}
```

## Available Hooks

### Authentication
- `useLogin()` - Login with email/password
- `useRegister()` - Register new business
- `useLogout()` - Logout current user
- `useCurrentUser()` - Get current authenticated user
- `useChangePassword()` - Change user password

### Users
- `useUsers(params?)` - Get all users
- `useUser(id)` - Get user by ID
- `useCreateUser()` - Create new user
- `useUpdateUser()` - Update user
- `useDeleteUser()` - Delete user

### Business & Branches
- `useBusiness()` - Get business info
- `useUpdateBusiness()` - Update business
- `useBranches(params?)` - Get all branches
- `useBranch(id)` - Get branch by ID
- `useCreateBranch()` - Create branch
- `useUpdateBranch()` - Update branch

### Products
- `useProducts(branchId, params?)` - Get products
- `useProduct(branchId, id)` - Get product by ID
- `useCreateProduct()` - Create product
- `useUpdateProduct()` - Update product
- `useCategories(branchId)` - Get categories
- `useProductVariants(branchId, productId)` - Get variants

### Sales
- `useSales(branchId, params?)` - Get sales
- `useSale(branchId, id)` - Get sale by ID
- `useCreateSale()` - Create sale
- `useRefundSale()` - Refund sale

### Inventory
- `useInventoryTransactions(branchId, params?)` - Get transactions
- `useAdjustStock()` - Adjust stock levels
- `useLowStockAlerts(branchId)` - Get low stock alerts

### Shifts
- `useShifts(branchId, params?)` - Get shifts
- `useActiveShift(branchId)` - Get currently active shift
- `useStartShift()` - Start new shift
- `useCloseShift()` - Close shift

### Customers
- `useCustomers(branchId, params?)` - Get customers
- `useCustomer(branchId, id)` - Get customer by ID
- `useCreateCustomer()` - Create customer
- `useSearchCustomerByPhone(branchId, phone)` - Search by phone

### Notifications
- `useAlerts(params?)` - Get alerts
- `useNotifications(params?)` - Get notifications
- `useUnreadAlertCount()` - Get unread count
- `useMarkAlertRead()` - Mark alert as read

### Reports
- `useSalesReport(branchId, params)` - Sales report
- `useInventoryReport(branchId)` - Inventory report
- `useRevenueReport(branchId, params)` - Revenue/profit report
- `useCustomerReport(branchId, params?)` - Customer analytics

### M-Pesa
- `useMpesaConfig()` - Get M-Pesa config
- `useUpdateMpesaConfig()` - Update M-Pesa config
- `useInitiateMpesaPayment()` - Initiate STK push

## Direct API Usage

You can also use the API functions directly without hooks:

```typescript
import { authApi, productApi, saleApi } from '@/frontend-client';

// Direct API calls
const user = await authApi.login({ email: '...', password: '...' });
const products = await productApi.getProducts(branchId);
const sale = await saleApi.createSale(branchId, saleData);
```

## Type Imports

```typescript
import type {
  User,
  Product,
  Sale,
  Customer,
  Branch,
  // ... and many more
} from '@/frontend-client';
```

## Authentication Flow

The client automatically handles authentication:

1. **Login** - Token is stored in localStorage
2. **Auto-attach** - Token is attached to all API requests
3. **Auto-logout** - On 401 responses, user is logged out and redirected to `/login`

## Error Handling

All hooks return error information through React Query:

```typescript
const { data, error, isLoading, isError } = useProducts(branchId);

if (isError) {
  console.error('Error:', error.message);
}
```

## Cache Management

React Query automatically manages caching. Mutations invalidate related queries:

```typescript
// Creating a product automatically refreshes the products list
const createProduct = useCreateProduct();
await createProduct.mutateAsync({ branchId, data: newProduct });
// useProducts(branchId) will automatically refetch
```

## License

MIT

## Support

For issues or questions, please contact the development team.
