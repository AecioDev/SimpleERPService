package models

import "time"

// CustomerDTO representa os dados de cliente para exibição
type CustomerDTO struct {
	ID             uint       `json:"id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	PersonType     string     `json:"person_type"`
	DocumentNumber string     `json:"document_number"`
	CompanyName    string     `json:"company_name"`
	IsActive       bool       `json:"is_active"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	Addresses      *[]Address `json:"addresses"`
	Contacts       *[]Contact `json:"contacts"`
}

// CustomerDetailDTO representa os dados detalhados de cliente para exibição
type CustomerDetailDTO struct {
	ID             uint   `json:"id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	PersonType     string `json:"person_type"`
	DocumentNumber string `json:"document_number"`
	CompanyName    string `json:"company_name"`
	IsActive       bool   `json:"is_active"`

	Addresses *[]Address `json:"addresses"`
	Contacts  *[]Contact `json:"contacts"`

	CreatedAt time.Time `json:"created_at"`
	CreatedBy *User     `json:"created_by"`

	UpdatedBy *User     `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`

	Sales        *[]Sale        `json:"sales"`
	Transactions *[]Transaction `json:"transactions"`
}

// CustomerListDTO representa uma lista paginada de clientes
type CustomerListDTO struct {
	Customer   []CustomerDTO `json:"data"`
	Pagination PaginationDTO `json:"pagination"`
}

// ToDTO converte um Customer para CustomerDTO
func (c *Customer) ToDTO() CustomerDTO {
	return CustomerDTO{
		ID:             c.ID,
		FirstName:      c.FirstName,
		LastName:       c.LastName,
		PersonType:     c.PersonType,
		DocumentNumber: c.DocumentNumber,
		CompanyName:    c.CompanyName,
		IsActive:       c.IsActive,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}

// ToDetailDTO converte um Customer para CustomerDetailDTO
func (c *Customer) ToDetailDTO() CustomerDetailDTO {
	return CustomerDetailDTO{
		ID:             c.ID,
		FirstName:      c.FirstName,
		LastName:       c.LastName,
		PersonType:     c.PersonType,
		DocumentNumber: c.DocumentNumber,
		CompanyName:    c.CompanyName,
		IsActive:       c.IsActive,
		CreatedAt:      c.CreatedAt,
		UpdatedAt:      c.UpdatedAt,
	}
}
