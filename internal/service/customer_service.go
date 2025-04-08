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
		Customer:   customerDTOs,
		Pagination: *models.ToPaginationDTO(pagination),
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
	document := strings.ReplaceAll(req.DocumentNumber, ".", "")
	document = strings.ReplaceAll(document, "-", "")
	document = strings.ReplaceAll(document, "/", "")

	// Criar cliente
	customer := models.Customer{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PersonType:     req.PersonType,
		DocumentNumber: document,
		CompanyName:    req.CompanyName,
		Notes:          req.Notes,
		CreatedByID:    &req.CreatedByID,
		IsActive:       true, // Por padrão, clientes são criados ativos
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

	// Atualizar campos básicos
	customer.FirstName = req.FirstName
	customer.LastName = req.LastName
	customer.CompanyName = req.CompanyName
	customer.IsActive = req.IsActive
	customer.Notes = req.Notes

	// Atualizar e formatar o DocumentNumber (removendo caracteres especiais)
	document := strings.NewReplacer(".", "", "-", "", "/", "").Replace(req.DocumentNumber)
	customer.DocumentNumber = document

	// TODO Implementar os Métodos
	//// Atualizar Relacionamentos (se IDs foram enviados)
	// if len(req.DocumentsIDs) > 0 {
	// 	documents, err := s.documentRepo.FindByIDs(req.DocumentsIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	customer.Documents = documents
	// }

	// if len(req.AdressesIDs) > 0 {
	// 	addresses, err := s.addressRepo.FindByIDs(req.AdressesIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	customer.Addresses = addresses
	// }

	// if len(req.ContactsIDs) > 0 {
	// 	contacts, err := s.contactRepo.FindByIDs(req.ContactsIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	customer.Contacts = contacts
	// }

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
