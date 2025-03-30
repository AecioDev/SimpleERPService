package service

import (
	"errors"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/utils"

	"gorm.io/gorm"
)

// UserService gerencia operações relacionadas a usuários
type UserService struct {
	db *gorm.DB
}

// NewUserService cria um novo serviço de usuários
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// GetUsers retorna uma lista paginada de usuários
func (s *UserService) GetUsers(pagination *utils.Pagination) ([]models.User, error) {
	var users []models.User

	query := s.db.Model(&models.User{}).Preload("Role")
	query, err := utils.Paginate(&models.User{}, pagination, query)
	if err != nil {
		return nil, err
	}

	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetUserByID busca um usuário pelo ID
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Role").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser cria um novo usuário
func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.User, error) {
	// Verificar se o username já existe
	var count int64
	s.db.Model(&models.User{}).Where("LOWER(username) = LOWER(?)", req.Username).Count(&count)
	if count > 0 {
		return nil, errors.New("username já está em uso")
	}

	// Verificar se o email já existe
	s.db.Model(&models.User{}).Where("LOWER(email) = LOWER(?)", req.Email).Count(&count)
	if count > 0 {
		return nil, errors.New("email já está em uso")
	}

	// Verificar se o perfil existe
	var role models.Role
	if err := s.db.First(&role, req.RoleID).Error; err != nil {
		return nil, errors.New("perfil não encontrado")
	}

	// Gerar hash da senha
	passwordHash, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Criar usuário
	user := models.User{
		Username:     req.Username,
		PasswordHash: passwordHash,
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		RoleID:       req.RoleID,
		IsActive:     true,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Carregar o perfil
	s.db.Preload("Role").First(&user, user.ID)

	return &user, nil
}

// UpdateUser atualiza um usuário existente
func (s *UserService) UpdateUser(id uint, req models.UpdateUserRequest) (*models.User, error) {
	// Buscar usuário
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	// Verificar se o email já está em uso por outro usuário
	if req.Email != "" && req.Email != user.Email {
		var count int64
		s.db.Model(&models.User{}).Where("LOWER(email) = LOWER(?) AND id != ?", req.Email, id).Count(&count)
		if count > 0 {
			return nil, errors.New("email já está em uso")
		}
	}

	// Verificar se o perfil existe
	if req.RoleID != 0 && req.RoleID != user.RoleID {
		var role models.Role
		if err := s.db.First(&role, req.RoleID).Error; err != nil {
			return nil, errors.New("perfil não encontrado")
		}
	}

	// Atualizar campos
	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.RoleID != 0 {
		user.RoleID = req.RoleID
	}

	// Salvar alterações
	if err := s.db.Save(&user).Error; err != nil {
		return nil, err
	}

	// Carregar o perfil
	s.db.Preload("Role").First(&user, user.ID)

	return &user, nil
}

// DeleteUser desativa um usuário
func (s *UserService) DeleteUser(id uint) error {
	// Buscar usuário
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	// Desativar usuário
	user.IsActive = false
	return s.db.Save(&user).Error
}

// ChangePassword altera a senha de um usuário
func (s *UserService) ChangePassword(id uint, currentPassword, newPassword string) error {
	// Buscar usuário
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return err
	}

	// Verificar senha atual
	if !utils.CheckPasswordHash(currentPassword, user.PasswordHash) {
		return errors.New("senha atual incorreta")
	}

	// Gerar hash da nova senha
	passwordHash, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// Atualizar senha
	user.PasswordHash = passwordHash
	return s.db.Save(&user).Error
}
