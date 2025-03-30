package service

import (
	"errors"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/utils"

	"gorm.io/gorm"
)

// RoleService gerencia operações relacionadas a perfis de usuário
type RoleService struct {
	db *gorm.DB
}

// NewRoleService cria um novo serviço de perfis
func NewRoleService(db *gorm.DB) *RoleService {
	return &RoleService{
		db: db,
	}
}

// GetRoles retorna uma lista paginada de perfis
func (s *RoleService) GetRoles(pagination *utils.Pagination) ([]models.Role, error) {
	var roles []models.Role

	query := s.db.Model(&models.Role{})
	query, err := utils.Paginate(&models.Role{}, pagination, query)
	if err != nil {
		return nil, err
	}

	if err := query.Find(&roles).Error; err != nil {
		return nil, err
	}

	return roles, nil
}

// GetRoleByID busca um perfil pelo ID
func (s *RoleService) GetRoleByID(id uint) (*models.Role, error) {
	var role models.Role
	if err := s.db.Preload("Permissions").First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// CreateRole cria um novo perfil
func (s *RoleService) CreateRole(req models.CreateRoleRequest) (*models.Role, error) {
	// Verificar se o nome já existe
	var count int64
	s.db.Model(&models.Role{}).Where("name = ?", req.Name).Count(&count)
	if count > 0 {
		return nil, errors.New("nome de perfil já está em uso")
	}

	// Criar perfil
	role := models.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.db.Create(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

// UpdateRole atualiza um perfil existente
func (s *RoleService) UpdateRole(id uint, req models.UpdateRoleRequest) (*models.Role, error) {
	// Buscar perfil
	var role models.Role
	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}

	// Verificar se o nome já está em uso por outro perfil
	if req.Name != "" && req.Name != role.Name {
		var count int64
		s.db.Model(&models.Role{}).Where("name = ? AND id != ?", req.Name, id).Count(&count)
		if count > 0 {
			return nil, errors.New("nome de perfil já está em uso")
		}
	}

	// Atualizar campos
	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Description != "" {
		role.Description = req.Description
	}

	// Salvar alterações
	if err := s.db.Save(&role).Error; err != nil {
		return nil, err
	}

	return &role, nil
}

// DeleteRole exclui um perfil
func (s *RoleService) DeleteRole(id uint) error {
	// Verificar se o perfil está sendo usado por usuários
	var count int64
	s.db.Model(&models.User{}).Where("role_id = ?", id).Count(&count)
	if count > 0 {
		return errors.New("não é possível excluir um perfil que está sendo usado por usuários")
	}

	// Excluir perfil
	return s.db.Delete(&models.Role{}, id).Error
}

// GetPermissions retorna todas as permissões
func (s *RoleService) GetPermissions() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := s.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

// GetPermissionsByModule retorna permissões agrupadas por módulo
func (s *RoleService) GetPermissionsByModule() ([]models.PermissionsByModule, error) {
	var permissions []models.Permission
	if err := s.db.Find(&permissions).Error; err != nil {
		return nil, err
	}

	// Agrupar permissões por módulo
	moduleMap := make(map[string][]models.Permission)
	for _, perm := range permissions {
		moduleMap[perm.Module] = append(moduleMap[perm.Module], perm)
	}

	// Converter mapa para slice
	var result []models.PermissionsByModule
	for module, perms := range moduleMap {
		result = append(result, models.PermissionsByModule{
			Module:      module,
			Permissions: perms,
		})
	}

	return result, nil
}

// UpdateRolePermissions atualiza as permissões de um perfil
func (s *RoleService) UpdateRolePermissions(id uint, permissionIDs []uint) (*models.Role, error) {
	// Buscar perfil
	var role models.Role
	if err := s.db.First(&role, id).Error; err != nil {
		return nil, err
	}

	// Buscar permissões
	var permissions []models.Permission
	if err := s.db.Where("id IN ?", permissionIDs).Find(&permissions).Error; err != nil {
		return nil, err
	}

	// Verificar se todas as permissões solicitadas existem
	if len(permissions) != len(permissionIDs) {
		return nil, errors.New("uma ou mais permissões não existem")
	}

	// Atualizar permissões do perfil
	if err := s.db.Model(&role).Association("Permissions").Replace(&permissions); err != nil {
		return nil, err
	}

	// Recarregar perfil com permissões
	if err := s.db.Preload("Permissions").First(&role, id).Error; err != nil {
		return nil, err
	}

	return &role, nil
}
