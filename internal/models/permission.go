package models

import (
	"time"

	"gorm.io/gorm"
)

// Permission representa uma permissão no sistema
type Permission struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:100;not null;unique" json:"name"`
	Description string         `json:"description"`
	Module      string         `gorm:"size:50;not null" json:"module"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// PermissionsByModule agrupa permissões por módulo
type PermissionsByModule struct {
	Module      string       `json:"module"`
	Permissions []Permission `json:"permissions"`
}
