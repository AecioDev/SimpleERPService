package service

import (
	"simple-erp-service/internal/models"
	"simple-erp-service/internal/repository"
	"simple-erp-service/internal/utils"
	"simple-erp-service/internal/validator"
	"strings"
)

// CustomerService gerencia operações relacionadas a clientes
type CustomerService struct {
	customerRepo repository.CustomerRepository
	validator    *validator.CustomerValidator
}

// NewCustomerService cria um novo serviço de clientes
func NewCustomerService(customerRepo repository.CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
		validator:    validator.NewCustomerValidator(customerRepo),
	}
}

// GetCustomers retorna uma lista paginada de clientes
func (s *CustomerService) GetCustomers(pagination *utils.Pagination) (*models.CustomerListDTO, error) {
	customers, err := s.customerRepo.FindAll(pagination)
	if err != nil {
		return nil, err
	}
	
	// Converter para DTOs
	customerDTOs := make([]models.CustomerDTO, 0, len(customers))
	for _, customer := range customers {
		customerDTOs = append(customerDTOs, customer.ToDTO())
	}
	
	return &models.CustomerListDTO{
		Data:       customerDTOs,
		Pagination: models.ToPaginationDTO(pagination),
	}, nil
}

// GetCustomerByID busca um cliente pelo ID
func (s *CustomerService) GetCustomerByID(id uint) (*models.CustomerDetailDTO, error) {
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, utils.ErrNotFound
	}
	
	// Converter para DTO
	customerDetailDTO := customer.ToDetailDTO()
	return &customerDetailDTO, nil
}

// CreateCustomer cria um novo cliente
func (s *CustomerService) CreateCustomer(req models.CreateCustomerRequest) (*models.CustomerDTO, error) {
	// Validar dados
	if err := s.validator.ValidateForCreation(req); err != nil {
		return nil, err
	}
	
	// Formatar documento (remover caracteres não numéricos)
	document := strings.ReplaceAll(req.Document, ".", "")
	document = strings.ReplaceAll(document, "-", "")
	document = strings.ReplaceAll(document, "/", "")
	
	// Criar cliente
	customer := models.Customer{
		Name:        req.Name,
		Document:    document,
		Email:       req.Email,
		Phone:       req.Phone,
		ContactName: req.ContactName,
		Address:     req.Address,
		City:        req.City,
		State:       req.State,
		ZipCode:     req.ZipCode,
		Notes:       req.Notes,
		IsActive:    true, // Por padrão, clientes são criados ativos
	}
	
	if err := s.customerRepo.Create(&customer); err != nil {
		return nil, err
	}
	
	// Converter para DTO
	customerDTO := customer.ToDTO()
	return &customerDTO, nil
}

// UpdateCustomer atualiza um cliente existente
func (s *CustomerService) UpdateCustomer(id uint, req models.UpdateCustomerRequest) (*models.CustomerDTO, error) {
	// Validar dados
	if err := s.validator.ValidateForUpdate(id, req); err != nil {
		return nil, err
	}
	
	// Buscar cliente
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if customer == nil {
		return nil, utils.ErrNotFound
	}
	
	// Atualizar campos
	if req.Name != "" {
		customer.Name = req.Name
	}
	if req.Document != "" {
		// Formatar documento (remover caracteres não numéricos)
		document := strings.ReplaceAll(req.Document, ".", "")
		document = strings.ReplaceAll(document, "-", "")
		document = strings.ReplaceAll(document, "/", "")
		customer.Document = document
	}
	if req.Email != "" {
		customer.Email = req.Email
	}
	if req.Phone != "" {
		customer.Phone = req.Phone
	}
	if req.ContactName != "" {
		customer.ContactName = req.ContactName
	}
	if req.Address != "" {
		customer.Address = req.Address
	}
	if req.City != "" {
		customer.City = req.City
	}
	if req.State != "" {
		customer.State = req.State
	}
	if req.ZipCode != "" {
		customer.ZipCode = req.ZipCode
	}
	if req.Notes != "" {
		customer.Notes = req.Notes
	}
	if req.IsActive != nil {
		customer.IsActive = *req.IsActive
	}
	
	// Salvar alterações
	if err := s.customerRepo.Update(customer); err != nil {
		return nil, err
	}
	
	// Converter para DTO
	customerDTO := customer.ToDTO()
	return &customerDTO, nil
}

// DeleteCustomer exclui um cliente (soft delete)
func (s *CustomerService) DeleteCustomer(id uint) error {
	// Verificar se o cliente existe
	customer, err := s.customerRepo.FindByID(id)
	if err != nil {
		return err
	}
	if customer == nil {
		return utils.ErrNotFound
	}
	
	// Excluir cliente
	return s.customerRepo.Delete(id)
}