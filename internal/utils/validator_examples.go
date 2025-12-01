package utils

// Example usage of custom validators

/*
Example 1: Simple validation

func (h *ProductHandler) Create(c *gin.Context) {
    var input struct {
        Name  string  `json:"name"`
        Price float64 `json:"price"`
        SKU   string  `json:"sku"`
    }

    if err := BindJSON(c, &input); err != nil {
        response.Error(c, 400, err.Error())
        return
    }

    validator := NewValidator()
    validator.
        Required("name", input.Name).
        MinLength("name", input.Name, 3).
        Required("sku", input.SKU).
        Min("price", input.Price, 0)

    if !validator.IsValid() {
        response.ValidationError(c, validator.Errors().ToMap())
        return
    }

    // Process valid input...
}

Example 2: Email validation

func (h *AuthHandler) Register(c *gin.Context) {
    var input struct {
        Email    string `json:"email"`
        Password string `json:"password"`
        Name     string `json:"name"`
    }

    if err := BindJSON(c, &input); err != nil {
        response.Error(c, 400, err.Error())
        return
    }

    validator := NewValidator()
    validator.
        Required("email", input.Email).
        Email("email", input.Email).
        Required("password", input.Password).
        MinLength("password", input.Password, 8).
        Required("name", input.Name)

    if !validator.IsValid() {
        response.ValidationError(c, validator.Errors().ToMap())
        return
    }

    // Process registration...
}

Example 3: OneOf validation (for enums)

func (h *SaleHandler) Create(c *gin.Context) {
    var input struct {
        PaymentMethod string  `json:"payment_method"`
        Total         float64 `json:"total"`
    }

    if err := BindJSON(c, &input); err != nil {
        response.Error(c, 400, err.Error())
        return
    }

    validator := NewValidator()
    validator.
        Required("payment_method", input.PaymentMethod).
        OneOf("payment_method", input.PaymentMethod, []string{
            constants.PaymentCash,
            constants.PaymentCard,
            constants.PaymentMPesa,
        }).
        Min("total", input.Total, 0)

    if !validator.IsValid() {
        response.ValidationError(c, validator.Errors().ToMap())
        return
    }

    // Process sale...
}

Example 4: Custom validation logic

func (h *UserHandler) Create(c *gin.Context) {
    var input struct {
        Role     string `json:"role"`
        BranchID string `json:"branch_id"`
    }

    if err := BindJSON(c, &input); err != nil {
        response.Error(c, 400, err.Error())
        return
    }

    validator := NewValidator()
    validator.
        Required("role", input.Role).
        OneOf("role", input.Role, []string{
            constants.RoleAdmin,
            constants.RoleManager,
            constants.RoleCashier,
        }).
        Custom("branch_id", "branch_id is required for non-admin roles",
            input.Role == constants.RoleAdmin || input.BranchID != "")

    if !validator.IsValid() {
        response.ValidationError(c, validator.Errors().ToMap())
        return
    }

    // Process user creation...
}

Example 5: UUID validation

func (h *ProductHandler) Update(c *gin.Context) {
    var input struct {
        CategoryID string  `json:"category_id"`
        Name       string  `json:"name"`
        Price      float64 `json:"price"`
    }

    if err := BindJSON(c, &input); err != nil {
        response.Error(c, 400, err.Error())
        return
    }

    validator := NewValidator()
    validator.
        UUID("category_id", input.CategoryID).
        Required("name", input.Name).
        Min("price", input.Price, 0)

    if !validator.IsValid() {
        response.ValidationError(c, validator.Errors().ToMap())
        return
    }

    // Process update...
}
*/
