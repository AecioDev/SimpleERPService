package service

import (
	"simple-erp-service/internal/models"
	"simple-erp-service/internal/repository"
	"simple-erp-service/internal/utils"
	"simple-erp-service/internal/validator"
	"strings"
)

// SupplierService gerencia operações relacionadas a fornecedores
type SupplierService struct {
	supplierRepo repository.SupplierRepository
	validator    *validator.SupplierValidator
}

// NewSupplierService cria um novo serviço de fornecedores
func NewSupplierService(supplierRepo repository.SupplierRepository) *SupplierService {
	return &SupplierService{
		supplierRepo: supplierRepo,
		validator:    validator.NewSupplierValidator(supplierRepo),
	}
}

// GetSuppliers retorna uma lista paginada de fornecedores
func (s *SupplierService) GetSuppliers(pagination *utils.Pagination) (*models.SupplierListDTO, error) {
	suppliers, err := s.supplierRepo.FindAll(pagination)
	if err != nil {
		return nil, err
	}

	// Converter para DTOs
	supplierDTOs := make([]models.SupplierDTO, 0, len(suppliers))
	for _, supplier := range suppliers {
		supplierDTOs = append(supplierDTOs, supplier.ToDTO())
	}

	return &models.SupplierListDTO{
		Supplier:   supplierDTOs,
		Pagination: *models.ToPaginationDTO(pagination),
	}, nil
}

// GetSupplierByID busca um fornecedor pelo ID
func (s *SupplierService) GetSupplierByID(id uint) (*models.SupplierDetailDTO, error) {
	supplier, err := s.supplierRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if supplier == nil {
		return nil, utils.ErrNotFound
	}

	// Converter para DTO
	supplierDetailDTO := supplier.ToDetailDTO()
	return &supplierDetailDTO, nil
}

// CreateSupplier cria um novo fornecedor
func (s *SupplierService) CreateSupplier(req models.CreateSupplierRequest) (*models.SupplierDTO, error) {
	// Validar dados
	if err := s.validator.ValidateForCreation(req); err != nil {
		return nil, err
	}

	// Formatar documento (remover caracteres não numéricos)
	document := strings.ReplaceAll(req.DocumentNumber, ".", "")
	document = strings.ReplaceAll(document, "-", "")
	document = strings.ReplaceAll(document, "/", "")

	// Criar fornecedor
	supplier := models.Supplier{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PersonType:     req.PersonType,
		DocumentNumber: document,
		CompanyName:    req.CompanyName,
		Notes:          req.Notes,
		CreatedByID:    &req.CreatedByID,
		IsActive:       true, // Por padrão, fornecedores são criados ativos
	}

	if err := s.supplierRepo.Create(&supplier); err != nil {
		return nil, err
	}

	// Converter para DTO
	supplierDTO := supplier.ToDTO()
	return &supplierDTO, nil
}

// UpdateSupplier atualiza um fornecedor existente
func (s *SupplierService) UpdateSupplier(id uint, req models.UpdateSupplierRequest) (*models.SupplierDTO, error) {
	// Validar dados
	if err := s.validator.ValidateForUpdate(id, req); err != nil {
		return nil, err
	}

	// Buscar fornecedor
	supplier, err := s.supplierRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if supplier == nil {
		return nil, utils.ErrNotFound
	}

	// Atualizar campos básicos
	supplier.FirstName = req.FirstName
	supplier.LastName = req.LastName
	supplier.CompanyName = req.CompanyName
	supplier.IsActive = req.IsActive
	supplier.Notes = req.Notes

	// Atualizar e formatar o DocumentNumber (removendo caracteres especiais)
	document := strings.NewReplacer(".", "", "-", "", "/", "").Replace(req.DocumentNumber)
	supplier.DocumentNumber = document

	// TODO Implementar os Métodos
	//// Atualizar Relacionamentos (se IDs foram enviados)
	// if len(req.DocumentsIDs) > 0 {
	// 	documents, err := s.documentRepo.FindByIDs(req.DocumentsIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	supplier.Documents = documents
	// }

	// if len(req.AdressesIDs) > 0 {
	// 	addresses, err := s.addressRepo.FindByIDs(req.AdressesIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	supplier.Addresses = addresses
	// }

	// if len(req.ContactsIDs) > 0 {
	// 	contacts, err := s.contactRepo.FindByIDs(req.ContactsIDs)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	supplier.Contacts = contacts
	// }

	// Salvar alterações
	if err := s.supplierRepo.Update(supplier); err != nil {
		return nil, err
	}

	// Converter para DTO
	supplierDTO := supplier.ToDTO()
	return &supplierDTO, nil
}

// DeleteSupplier exclui um fornecedor (soft delete)
func (s *SupplierService) DeleteSupplier(id uint) error {
	// Verificar se o fornecedor existe
	supplier, err := s.supplierRepo.FindByID(id)
	if err != nil {
		return err
	}
	if supplier == nil {
		return utils.ErrNotFound
	}

	// Excluir fornecedor
	return s.supplierRepo.Delete(id)
}
