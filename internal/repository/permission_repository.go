package repository

import (
	"simple-erp-service/internal/models"

	"gorm.io/gorm"
)

// PermissionRepository define as operações de acesso a dados para permissões
type PermissionRepository interface {
	Repository
	FindAll() ([]models.Permission, error)
	FindByIDs(ids []uint) ([]models.Permission, error)
	GroupByModule() (map[string][]models.Permission, error)
}

// GormPermissionRepository implementa PermissionRepository usando GORM
type GormPermissionRepository struct {
	*BaseRepository
}

// NewPermissionRepository cria um novo repository de permissões
func NewPermissionRepository(db *gorm.DB) PermissionRepository {
	return &GormPermissionRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// FindAll retorna todas as permissões
func (r *GormPermissionRepository) FindAll() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.GetDB().Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

// FindByIDs busca permissões pelos IDs
func (r *GormPermissionRepository) FindByIDs(ids []uint) ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.GetDB().Where("id IN ?", ids).Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

// GroupByModule retorna permissões agrupadas por módulo
func (r *GormPermissionRepository) GroupByModule() (map[string][]models.Permission, error) {
	var permissions []models.Permission
	if err := r.GetDB().Find(&permissions).Error; err != nil {
		return nil, err
	}

	// Agrupar permissões por módulo
	moduleMap := make(map[string][]models.Permission)
	for _, perm := range permissions {
		moduleMap[perm.Module] = append(moduleMap[perm.Module], perm)
	}

	return moduleMap, nil
}
