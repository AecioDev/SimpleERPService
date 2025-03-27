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

// UserHandler gerencia as requisições relacionadas a usuários
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler cria um novo handler de usuários
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		userService: service.NewUserService(db),
	}
}

// GetUsers retorna uma lista paginada de usuários
func (h *UserHandler) GetUsers(c *gin.Context) {
	pagination := utils.GetPaginationParams(c)

	users, err := h.userService.GetUsers(&pagination)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Erro ao buscar usuários", err.Error())
		return
	}

	// Converter para resposta
	var response []models.UserResponse
	for _, user := range users {
		response = append(response, user.ToResponse())
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuários encontrados", response, pagination)
}

// GetUser retorna um usuário específico
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		utils.ErrorResponse(c, http.StatusNotFound, "Usuário não encontrado", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuário encontrado", user.ToResponse(), nil)
}

// CreateUser cria um novo usuário
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao criar usuário", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, "Usuário criado com sucesso", user.ToResponse(), nil)
}

// UpdateUser atualiza um usuário existente
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	user, err := h.userService.UpdateUser(uint(id), req)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao atualizar usuário", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuário atualizado com sucesso", user.ToResponse(), nil)
}

// DeleteUser desativa um usuário
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	if err := h.userService.DeleteUser(uint(id)); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao desativar usuário", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Usuário desativado com sucesso", nil, nil)
}

// ChangePassword altera a senha de um usuário
func (h *UserHandler) ChangePassword(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "ID inválido", err.Error())
		return
	}

	// Verificar se o usuário está alterando sua própria senha
	userID, exists := c.Get("userID")
	if !exists || userID.(uint) != uint(id) {
		// Verificar se o usuário tem permissão para alterar senhas de outros usuários
		role, exists := c.Get("role")
		if !exists || role.(string) != "ADMIN" {
			utils.ErrorResponse(c, http.StatusForbidden, "Acesso negado", "Você não tem permissão para alterar a senha de outro usuário")
			return
		}
	}

	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ValidationErrorResponse(c, "Dados inválidos", err.Error())
		return
	}

	if err := h.userService.ChangePassword(uint(id), req.CurrentPassword, req.NewPassword); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "Erro ao alterar senha", err.Error())
		return
	}

	utils.SuccessResponse(c, http.StatusOK, "Senha alterada com sucesso", nil, nil)
}
