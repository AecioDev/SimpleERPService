package models

import "time"

// SupplierDTO representa os dados de fornecedor para exibição
type SupplierDTO struct {
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

// SupplierDetailDTO representa os dados detalhados de fornecedor para exibição
type SupplierDetailDTO struct {
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
	BankInfo     string    `json:"bank_info,omitempty"`
	Notes        string    `json:"notes,omitempty"`
	IsActive     bool      `json:"is_active"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// SupplierListDTO representa uma lista paginada de fornecedores
type SupplierListDTO struct {
	Data       []SupplierDTO  `json:"data"`
	Pagination PaginationDTO  `json:"pagination"`
}

// ToDTO converte um Supplier para SupplierDTO
func (s *Supplier) ToDTO() SupplierDTO {
	return SupplierDTO{
		ID:          s.ID,
		Name:        s.Name,
		Document:    s.Document,
		Email:       s.Email,
		Phone:       s.Phone,
		ContactName: s.ContactName,
		IsActive:    s.IsActive,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}

// ToDetailDTO converte um Supplier para SupplierDetailDTO
func (s *Supplier) ToDetailDTO() SupplierDetailDTO {
	return SupplierDetailDTO{
		ID:          s.ID,
		Name:        s.Name,
		Document:    s.Document,
		Email:       s.Email,
		Phone:       s.Phone,
		ContactName: s.ContactName,
		Address:     s.Address,
		City:        s.City,
		State:       s.State,
		ZipCode:     s.ZipCode,
		BankInfo:    s.BankInfo,
		Notes:       s.Notes,
		IsActive:    s.IsActive,
		CreatedAt:   s.CreatedAt,
		UpdatedAt:   s.UpdatedAt,
	}
}