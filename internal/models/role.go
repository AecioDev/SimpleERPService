package models

import (
	"time"

	"gorm.io/gorm"
)

// Role representa um perfil de usuário
type Role struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"size:50;not null;unique" json:"name"`
	Description string         `json:"description"`
	Permissions []*Permission  `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// CreateRoleRequest representa os dados para criar um novo perfil
type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

// UpdateRoleRequest representa os dados para atualizar um perfil
type UpdateRoleRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// UpdateRolePermissionsRequest representa os dados para atualizar permissões de um perfil
type UpdateRolePermissionsRequest struct {
	PermissionIDs []uint `json:"permission_ids" binding:"required"`
}
