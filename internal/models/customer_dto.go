package models

import "time"

// CustomerDTO representa os dados de cliente para exibição
type CustomerDTO struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Document     string    `json:"document"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	ContactName  string    `json:"contact_name,omitempty"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CustomerDetailDTO representa os dados detalhados de cliente para exibição
type CustomerDetailDTO struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Document     string    `json:"document"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	ContactName  string    `json:"contact_name,omitempty"`
	Address      string    `json:"address,omitempty"`
	City         string    `json:"city,omitempty"`
	State        string    `json:"state,omitempty"`
	ZipCode      string    `json:"zip_code,omitempty"`
	Notes        string    `json:"notes,omitempty"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// CustomerListDTO representa uma lista paginada de clientes
type CustomerListDTO struct {
	Data       []CustomerDTO  `json:"data"`
	Pagination PaginationDTO  `json:"pagination"`
}

// ToDTO converte um Customer para CustomerDTO
func (c *Customer) ToDTO() CustomerDTO {
	return CustomerDTO{
		ID:          c.ID,
		Name:        c.Name,
		Document:    c.Document,
		Email:       c.Email,
		Phone:       c.Phone,
		ContactName: c.ContactName,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

// ToDetailDTO converte um Customer para CustomerDetailDTO
func (c *Customer) ToDetailDTO() CustomerDetailDTO {
	return CustomerDetailDTO{
		ID:          c.ID,
		Name:        c.Name,
		Document:    c.Document,
		Email:       c.Email,
		Phone:       c.Phone,
		ContactName: c.ContactName,
		Address:     c.Address,
		City:        c.City,
		State:       c.State,
		ZipCode:     c.ZipCode,
		Notes:       c.Notes,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}