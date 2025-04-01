package service_test

import (
	"errors"
	"testing"

	"simple-erp-service/internal/models"
	"simple-erp-service/internal/repository"
	"simple-erp-service/internal/service"
	"simple-erp-service/internal/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetDB() *gorm.DB {
	args := m.Called()
	return args.Get(0).(*gorm.DB)
}

func (m *MockUserRepository) WithContext(ctx context.Context) repository.Repository {
	args := m.Called(ctx)
	return args.Get(0).(repository.Repository)
}

func (m *MockUserRepository) WithTx(tx *gorm.DB) repository.Repository {
	args := m.Called(tx)
	return args.Get(0).(repository.Repository)
}

func (m *MockUserRepository) FindAll(pagination *utils.Pagination) ([]models.User, error) {
	args := m.Called(pagination)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByIDWithRole(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByUsername(username string) (*models.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Update(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserRepository) ExistsByUsername(username string) (bool, error) {
	args := m.Called(username)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) ExistsByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) ExistsByUsernameExcept(username string, id uint) (bool, error) {
	args := m.Called(username, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) ExistsByEmailExcept(email string, id uint) (bool, error) {
	args := m.Called(email, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) CountByRoleID(roleID uint) (int64, error) {
	args := m.Called(roleID)
	return args.Get(0).(int64), args.Error(1)
}

// Mock do RoleRepository
type MockRoleRepository struct {
	mock.Mock
}

// Implementar métodos do RoleRepository...

// Testes do UserService
func TestGetUserByID(t *testing.T) {
	// Arrange
	mockUserRepo := new(MockUserRepository)
	mockRoleRepo := new(MockRoleRepository)
	
	userService := service.NewUserService(mockUserRepo, mockRoleRepo)
	
	testCases := []struct {
		name          string
		userID        uint
		mockSetup     func()
		expectedError error
		expectedUser  *models.UserDetailDTO
	}{
		{
			name:   "Success",
			userID: 1,
			mockSetup: func() {
				user := &models.User{
					Base: models.Base{ID: 1},
					Username: "testuser",
					Name: "Test User",
					Email: "test@example.com",
					RoleID: 1,
					Role: models.Role{
						Base: models.Base{ID: 1},
						Name: "Admin",
					},
					IsActive: true,
				}
				mockUserRepo.On("FindByIDWithRole", uint(1)).Return(user, nil)
			},
			expectedError: nil,
			expectedUser: &models.UserDetailDTO{
				ID: 1,
				Username: "testuser",
				Name: "Test User",
				Email: "test@example.com",
				RoleID: 1,
				Role: models.RoleDTO{
					ID: 1,
					Name: "Admin",
				},
				IsActive: true,
			},
		},
		{
			name:   "User Not Found",
			userID: 999,
			mockSetup: func() {
				mockUserRepo.On("FindByIDWithRole", uint(999)).Return(nil, nil)
			},
			expectedError: utils.ErrNotFound,
			expectedUser:  nil,
		},
		{
			name:   "Database Error",
			userID: 2,
			mockSetup: func() {
				mockUserRepo.On("FindByIDWithRole", uint(2)).Return(nil, errors.New("database error"))
			},
			expectedError: errors.New("database error"),
			expectedUser:  nil,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Reset mocks
			mockUserRepo = new(MockUserRepository)
			mockRoleRepo = new(MockRoleRepository)
			userService = service.NewUserService(mockUserRepo, mockRoleRepo)
			
			// Setup mocks
			tc.mockSetup()
			
			// Act
			user, err := userService.GetUserByID(tc.userID)
			
			// Assert
			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedUser.ID, user.ID)
				assert.Equal(t, tc.expectedUser.Username, user.Username)
				assert.Equal(t, tc.expectedUser.Name, user.Name)
				assert.Equal(t, tc.expectedUser.Email, user.Email)
				assert.Equal(t, tc.expectedUser.RoleID, user.RoleID)
				assert.Equal(t, tc.expectedUser.Role.ID, user.Role.ID)
				assert.Equal(t, tc.expectedUser.Role.Name, user.Role.Name)
				assert.Equal(t, tc.expectedUser.IsActive, user.IsActive)
			}
			
			// Verify expectations
			mockUserRepo.AssertExpectations(t)
			mockRoleRepo.AssertExpectations(t)
		})
	}
}

// Adicionar mais testes para outros métodos do UserService...