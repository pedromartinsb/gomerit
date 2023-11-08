package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateOpening
type CreateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" && r.Company == "" && r.Location == "" && r.Link == "" && r.Remote == nil && r.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}

// UpdateOpening
type UpdateOpeningRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningRequest) Validate() error {
	// If any field is provided, validation is truthy
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Link != "" || r.Remote != nil || r.Salary > 0 {
		return nil
	}
	// If none of the fields were provided, then return falsy
	return fmt.Errorf("at least one valid field must be provided")
}

// CreateHolding
type CreateHoldingRequest struct {
	Name string `json:"name"`
}

func (r *CreateHoldingRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	return nil
}

// UpdateHolding
type UpdateHoldingRequest struct {
	Name string `json:"name"`
}

func (r *UpdateHoldingRequest) Validate() error {
	if r.Name != "" {
		return nil
	}
	return fmt.Errorf("valid name filed must be provided")
}

// CreateCompany
type CreateCompanyRequest struct {
	Name      string `json:"name"`
	HoldingID uint   `json:"holding_id"`
}

func (r *CreateCompanyRequest) Validate() error {
	if r.Name == "" {
		return errParamIsRequired("name", "string")
	}
	if r.HoldingID == 0 {
		return errParamIsRequired("holding_id", "uint")
	}
	return nil
}
