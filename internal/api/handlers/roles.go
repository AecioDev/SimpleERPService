package handlers

import (
	"net/http"
	"strconv"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/service"
	"simple-erp-service/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RoleHandler gerencia as requisições relacionadas a perfis
type RoleHandler struct {
	roleService *service.RoleService
}

// NewRoleHandler cria um novo handler de perfis
func NewRoleHandler(db *gorm.DB) *RoleHandler {
	return &RoleHandler{
		roleService: service.NewRoleService(db),
	}
}

// GetRoles retorna uma lista paginada de perfis
func (h *RoleHandler) GetRoles(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)

	roles, err := h.roleService.GetRoles(&pagination)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao buscar perfis", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Perfis encontrados", roles, pagination)
}

// GetRole retorna um perfil específico
func (h *RoleHandler) GetRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	role, err := h.roleService.GetRoleByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Perfil não encontrado", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Perfil encontrado", role, nil)
}

// CreateRole cria um novo perfil
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req models.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	role, err := h.roleService.CreateRole(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao criar perfil", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Perfil criado com sucesso", role, nil)
}

// UpdateRole atualiza um perfil existente
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	var req models.UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	role, err := h.roleService.UpdateRole(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao atualizar perfil", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Perfil atualizado com sucesso", role, nil)
}

// DeleteRole exclui um perfil
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	if err := h.roleService.DeleteRole(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao excluir perfil", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Perfil excluído com sucesso", nil, nil)
}

// GetPermissions retorna todas as permissões
func (h *RoleHandler) GetPermissions(c *gin.Context) {
	permissions, err := h.roleService.GetPermissions()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao buscar permissões", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Permissões encontradas", permissions, nil)
}

// GetPermissionsByModule retorna permissões agrupadas por módulo
func (h *RoleHandler) GetPermissionsByModule(c *gin.Context) {
	permissions, err := h.roleService.GetPermissionsByModule()
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao buscar permissões", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Permissões encontradas", permissions, nil)
}

// UpdateRolePermissions atualiza as permissões de um perfil
func (h *RoleHandler) UpdateRolePermissions(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	var req models.UpdateRolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	role, err := h.roleService.UpdateRolePermissions(uint(id), req.PermissionIDs)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao atualizar permissões", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Permissões atualizadas com sucesso", role, nil)
}
