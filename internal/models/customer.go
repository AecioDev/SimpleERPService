package models

import (
	"time"

	"gorm.io/gorm"
)

// Customer representa um cliente
type Customer struct {
	gorm.Model

	FirstName      string `gorm:"size:100;" json:"first_name"`           // Nome
	LastName       string `gorm:"size:100;" json:"last_name"`            // Nome
	PersonType     string `gorm:"default:F" json:"person_type"`          // Pessoa Jurídica ou Física
	DocumentNumber string `gorm:"size:20;unique" json:"document_number"` // CPF ou CNPJ pra facilitar a busca aqui
	CompanyName    string `gorm:"size:100;" json:"company_name"`         // Razão Social
	IsActive       bool   `gorm:"default:true" json:"is_active"`         // Status do Cliente

	// Usuário que criou
	CreatedByID *uint `gorm:"column:created_by" json:"created_by"`
	CreatedBy   *User `gorm:"foreignKey:CreatedByID" json:"created_by_user,omitempty"`

	// Usuário que Atualizou
	UpdatedByID *uint `gorm:"column:updated_by" json:"updated_by"`
	UpdatedBy   *User `gorm:"foreignKey:UpdatedByID" json:"updated_by_user,omitempty"`

	Notes     string `gorm:"size:255" json:"notes"` // Observações
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relacionamentos 1...N 1 Cliente pra muitos ...
	Documents    []Document    `gorm:"foreignKey:CustomerID" json:"-"` // rg, inscrição estadual, inscrição municipal, cnh etc.
	Addresses    []Address     `gorm:"foreignKey:CustomerID" json:"-"` // Endereços do Cliente
	Contacts     []Contact     `gorm:"foreignKey:CustomerID" json:"-"` // email, telefones, etc.
	Sales        []Sale        `gorm:"foreignKey:CustomerID" json:"-"` // Vendas associadas ao cliente
	Transactions []Transaction `gorm:"foreignKey:CustomerID" json:"-"` // Transações associadas ao cliente (Titulos gerados no financeiro)

	// Outros relacionamentos podem ser adicionados conforme necessário
	// Exemplo: Invoices, Receipts, etc.

}

// TableName especifica o nome da tabela
func (Customer) TableName() string {
	return "customers"
}

// CreateCustomerRequest representa os dados para criar um novo cliente
// Especificar no Front um Estilo Passo a Passo para Gravar o Cliente no Primeiro Passo, e os Demais serem Update com os Demais dados.
type CreateCustomerRequest struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"omitempty"`
	PersonType     string `json:"person_type" binding:"required"` // Solicitar Atenção no Front pq não poderá alterar mais depois
	DocumentNumber string `json:"document_number" binding:"required"`
	CompanyName    string `json:"company_name" binding:"omitempty"`
	IsActive       bool   `json:"is_active" binding:"required"`
	Notes          string `json:"notes"`
	CreatedByID    uint   `json:"created_by" binding:"required"`
}

// UpdateCustomerRequest representa os dados para atualizar um cliente
type UpdateCustomerRequest struct {
	FirstName      string `json:"first_name" binding:"required"`
	LastName       string `json:"last_name" binding:"omitempty"`
	DocumentNumber string `json:"document_number" binding:"required"`
	CompanyName    string `json:"company_name" binding:"omitempty"`
	IsActive       bool   `json:"is_active" binding:"required"`
	Notes          string `json:"notes"`

	// Relacionamentos O Front manda uma Lista de IDs válidos para cada entidade relacionada.
	DocumentsIDs []uint `json:"documents_ids" binding:"omitempty"`
	AdressesIDs  []uint `json:"adresses_ids" binding:"omitempty"`
	ContactsIDs  []uint `json:"contacts_ids" binding:"omitempty"`
}
