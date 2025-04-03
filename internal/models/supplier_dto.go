package models

import "time"

// SupplierDTO representa os dados de fornecedor para exibição
type SupplierDTO struct {
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

// SupplierDetailDTO representa os dados detalhados de fornecedor para exibição
type SupplierDetailDTO struct {
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

	Purchases    *[]Purchase    `json:"purchases"`
	Transactions *[]Transaction `json:"transactions"`
}

// SupplierListDTO representa uma lista paginada de fornecedores
type SupplierListDTO struct {
	Supplier   []SupplierDTO `json:"data"`
	Pagination PaginationDTO `json:"pagination"`
}

// ToDTO converte um Supplier para SupplierDTO
func (s *Supplier) ToDTO() SupplierDTO {
	return SupplierDTO{
		ID:             s.ID,
		FirstName:      s.FirstName,
		LastName:       s.LastName,
		PersonType:     s.PersonType,
		DocumentNumber: s.DocumentNumber,
		CompanyName:    s.CompanyName,
		IsActive:       s.IsActive,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}

// ToDetailDTO converte um Supplier para SupplierDetailDTO
func (s *Supplier) ToDetailDTO() SupplierDetailDTO {
	return SupplierDetailDTO{
		ID:             s.ID,
		FirstName:      s.FirstName,
		LastName:       s.LastName,
		PersonType:     s.PersonType,
		DocumentNumber: s.DocumentNumber,
		CompanyName:    s.CompanyName,
		IsActive:       s.IsActive,
		CreatedAt:      s.CreatedAt,
		UpdatedAt:      s.UpdatedAt,
	}
}
