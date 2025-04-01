package service

import (
	"simple-erp-service/internal/models"
	"simple-erp-service/internal/repository"
	"simple-erp-service/internal/utils"
	"simple-erp-service/internal/validator"
)

// RoleService gerencia operações relacionadas a perfis de usuário
type RoleService struct {
	roleRepo  repository.RoleRepository
	userRepo  repository.UserRepository
	permRepo  repository.PermissionRepository
	validator *validator.RoleValidator
}

// NewRoleService cria um novo serviço de perfis
func NewRoleService(
	roleRepo repository.RoleRepository,
	userRepo repository.UserRepository,
	permRepo repository.PermissionRepository,
) *RoleService {
	return &RoleService{
		roleRepo:  roleRepo,
		userRepo:  userRepo,
		permRepo:  permRepo,
		validator: validator.NewRoleValidator(roleRepo, userRepo, permRepo),
	}
}

// GetRoles retorna uma lista paginada de perfis
func (s *RoleService) GetRoles(pagination *utils.Pagination) ([]models.RoleDTO, error) {
	roles, err := s.roleRepo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	// Converter para DTOs
	roleDTOs := make([]models.RoleDTO, 0, len(roles))
	for _, role := range roles {
		roleDTOs = append(roleDTOs, role.ToDTO())
	}

	return roleDTOs, nil
}

// GetRoleByID busca um perfil pelo ID
func (s *RoleService) GetRoleByID(id uint) (*models.RoleDetailDTO, error) {
	role, err := s.roleRepo.FindByIDWithPermissions(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, utils.ErrNotFound
	}

	// Converter para DTO
	roleDetailDTO := role.ToDetailDTO()
	return &roleDetailDTO, nil
}

// CreateRole cria um novo perfil
func (s *RoleService) CreateRole(req models.CreateRoleRequest) (*models.RoleDTO, error) {
	// Validar dados
	if err := s.validator.ValidateForCreation(req); err != nil {
		return nil, err
	}

	// Criar perfil
	role := models.Role{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.roleRepo.Create(&role); err != nil {
		return nil, err
	}

	// Converter para DTO
	roleDTO := role.ToDTO()
	return &roleDTO, nil
}

// UpdateRole atualiza um perfil existente
func (s *RoleService) UpdateRole(id uint, req models.UpdateRoleRequest) (*models.RoleDTO, error) {
	// Validar dados
	if err := s.validator.ValidateForUpdate(id, req); err != nil {
		return nil, err
	}

	// Buscar perfil
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, utils.ErrNotFound
	}

	// Atualizar campos
	if req.Name != "" {
		role.Name = req.Name
	}
	if req.Description != "" {
		role.Description = req.Description
	}

	// Salvar alterações
	if err := s.roleRepo.Update(role); err != nil {
		return nil, err
	}

	// Converter para DTO
	roleDTO := role.ToDTO()
	return &roleDTO, nil
}

// DeleteRole exclui um perfil
func (s *RoleService) DeleteRole(id uint) error {
	// Validar se o perfil pode ser excluído
	if err := s.validator.ValidateForDeletion(id); err != nil {
		return err
	}

	// Excluir perfil
	return s.roleRepo.Delete(id)
}

// GetPermissions retorna todas as permissões
func (s *RoleService) GetPermissions() ([]models.PermissionDTO, error) {
	permissions, err := s.permRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// Converter para DTOs
	permissionDTOs := make([]models.PermissionDTO, 0, len(permissions))
	for _, perm := range permissions {
		permissionDTOs = append(permissionDTOs, perm.ToDTO())
	}

	return permissionDTOs, nil
}

// GetPermissionsByModule retorna permissões agrupadas por módulo
func (s *RoleService) GetPermissionsByModule() ([]models.PermissionsByModuleDTO, error) {
	moduleMap, err := s.permRepo.GroupByModule()
	if err != nil {
		return nil, err
	}

	// Converter mapa para slice
	var result []models.PermissionsByModule
	for module, perms := range moduleMap {
		result = append(result, models.PermissionsByModule{
			Module:      module,
			Permissions: perms,
		})
	}

	// Converter para DTOs
	resultDTOs := make([]models.PermissionsByModuleDTO, 0, len(result))
	for _, item := range result {
		resultDTOs = append(resultDTOs, item.ToDTO())
	}

	return resultDTOs, nil
}

// UpdateRolePermissions atualiza as permissões de um perfil
func (s *RoleService) UpdateRolePermissions(id uint, permissionIDs []uint) (*models.RoleDetailDTO, error) {
	// Validar dados
	if err := s.validator.ValidatePermissionUpdate(id, permissionIDs); err != nil {
		return nil, err
	}

	// Buscar perfil
	role, err := s.roleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, utils.ErrNotFound
	}

	// Atualizar permissões
	if err := s.roleRepo.UpdatePermissions(role, permissionIDs); err != nil {
		return nil, err
	}

	// Buscar perfil atualizado com permissões
	updatedRole, err := s.roleRepo.FindByIDWithPermissions(id)
	if err != nil {
		return nil, err
	}

	// Converter para DTO
	roleDetailDTO := updatedRole.ToDetailDTO()
	return &roleDetailDTO, nil
}
