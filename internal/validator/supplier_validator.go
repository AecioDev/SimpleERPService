package validator

import (
	"simple-erp-service/internal/models"
	"simple-erp-service/internal/repository"
	"strings"
)

// SupplierValidator valida regras de negócio relacionadas a fornecedores
type SupplierValidator struct {
	supplierRepo repository.SupplierRepository
}

// NewSupplierValidator cria um novo validador de fornecedores
func NewSupplierValidator(supplierRepo repository.SupplierRepository) *SupplierValidator {
	return &SupplierValidator{
		supplierRepo: supplierRepo,
	}
}

// ValidateForCreation valida os dados para criação de um fornecedor
func (v *SupplierValidator) ValidateForCreation(req models.CreateSupplierRequest) error {
	var errors ValidationErrors

	// Verificar se o documento já existe
	if req.Document != "" {
		// Remover caracteres não numéricos para comparação
		document := strings.ReplaceAll(req.Document, ".", "")
		document = strings.ReplaceAll(document, "-", "")
		document = strings.ReplaceAll(document, "/", "")
		
		exists, err := v.supplierRepo.ExistsByDocument(document)
		if err != nil {
			return err
		}
		if exists {
			errors.AddError("document", "documento já está em uso")
		}
	}

	// Verificar se o email já existe (se fornecido)
	if req.Email != "" {
		exists, err := v.supplierRepo.ExistsByEmail(req.Email)
		if err != nil {
			return err
		}
		if exists {
			errors.AddError("email", "email já está em uso")
		}
	}

	// Validar nome
	if len(req.Name) < 3 {
		errors.AddError("name", "nome deve ter pelo menos 3 caracteres")
	}

	if errors.HasErrors() {
		return errors
	}
	return nil
}

// ValidateForUpdate valida os dados para atualização de um fornecedor
func (v *SupplierValidator) ValidateForUpdate(id uint, req models.UpdateSupplierRequest) error {
	var errors ValidationErrors

	// Verificar se o fornecedor existe
	supplier, err := v.supplierRepo.FindByID(id)
	if err != nil {
		return err
	}
	if supplier == nil {
		errors.AddError("id", "fornecedor não encontrado")
		return errors
	}

	// Verificar se o documento já está em uso por outro fornecedor (se fornecido)
	if req.Document != "" && req.Document != supplier.Document {
		// Remover caracteres não numéricos para comparação
		document := strings.ReplaceAll(req.Document, ".", "")
		document = strings.ReplaceAll(document, "-", "")
		document = strings.ReplaceAll(document, "/", "")
		
		exists, err := v.supplierRepo.ExistsByDocumentExcept(document, id)
		if err != nil {
			return err
		}
		if exists {
			errors.AddError("document", "documento já está em uso")
		}
	}

	// Verificar se o email já está em uso por outro fornecedor (se fornecido)
	if req.Email != "" && req.Email != supplier.Email {
		exists, err := v.supplierRepo.ExistsByEmailExcept(req.Email, id)
		if err != nil {
			return err
		}
		if exists {
			errors.AddError("email", "email já está em uso")
		}
	}

	// Validar nome (se fornecido)
	if req.Name != "" && len(req.Name) < 3 {
		errors.AddError("name", "nome deve ter pelo menos 3 caracteres")
	}

	if errors.HasErrors() {
		return errors
	}
	return nil
}