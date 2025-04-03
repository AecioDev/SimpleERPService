package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model

	Contact     string     `gorm:"unique;not null" json:"contact"` // Contato em si o email, o telefone etc
	ContactType string     `gorm:"not null" json:"contact_type"`   // Tipo de Contato: email, Telefone, Celular
	ContactName *time.Time `json:"contact_name"`                   // Um Nome para contato se for o caso
}

// TableName especifica o nome da tabela
func (Contact) TableName() string {
	return "contact"
}
