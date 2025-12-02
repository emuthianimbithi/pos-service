# API Client - Backend Response Handling

## Backend Response Format

The backend uses a custom response wrapper defined in `pkg/response/response.go`:

```go
type Response struct {
    Success bool        `json:"success"`
    Message string      `json:"message,omitempty"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
    Meta    interface{} `json:"meta,omitempty"`
}
```

### Success Response Example
```json
{
  "success": true,
  "message": "User created successfully",
  "data": {
    "id": "uuid",
    "email": "user@example.com",
    ...
  }
}
```

### Error Response Example
```json
{
  "success": false,
  "error": "Invalid credentials"
}
```

### Validation Error Example
```json
{
  "success": false,
  "error": "Validation failed",
  "details": {
    "email": "Email is required",
    "password": "Password must be at least 8 characters"
  }
}
```

### Paginated Response Example
```json
{
  "success": true,
  "data": [...],
  "meta": {
    "current_page": 1,
    "per_page": 20,
    "total": 100,
    "total_pages": 5
  }
}
```

## How the Client Handles This

The API client's response interceptor handles the unwrapping automatically:

1. **Success responses**: The `data` field from the backend response is returned
2. **Error responses**: The `error` and `details` fields are extracted and thrown
3. **Validation errors**: The `details` object is preserved in the error

## Usage in API Functions

All API functions should access `response.data.data` since:
- `response.data` = the full backend response `{ success, data, message }`
- `response.data.data` = the actual payload we want

Example:
```typescript
export const authApi = {
  login: async (credentials: LoginRequest): Promise<LoginResponse> => {
    const response = await apiClient.post<BackendResponse<LoginResponse>>('/auth/login', credentials);
    // response.data = { success: true, data: { token, user }, message: "..." }
    // response.data.data = { token, user }
    return response.data.data!;
  },
};
```

## Error Handling

Errors are automatically formatted and include validation details:

```typescript
try {
  await authApi.login(credentials);
} catch (error) {
  console.error(error.error); // Main error message
  console.error(error.details); // Validation errors if any
}
```
