package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	gorm.Model

	DocumentType         string     `gorm:"not null" json:"document_type"`          // Tipo de Documento: CPF, RG, CNPJ, IE, IM, CNH etc.
	DocumentOwner        string     `gorm:"not null" json:"document_owner"`         // Cliente ou Fornecedor etc.
	DocumentNumber       string     `gorm:"unique;not null" json:"document_number"` // Número do Documento
	DocumentValidate     *time.Time `json:"document_validate"`                      // Data de Validade (nullable)
	DocumentEmissionDate *time.Time `json:"document_emission_date"`                 // Data de Emissão (nullable)
	DocumentDepartment   string     `json:"document_department"`                    // Órgão Emissor
	StateID              uint       `json:"state_id"`                               // FK para State (UF de Emissão)
	State                State      `gorm:"foreignKey:StateID" json:"state"`        // Associação com a UF
}

// TableName especifica o nome da tabela
func (Document) TableName() string {
	return "document"
}
