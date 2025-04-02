package validator

import (
	"simple-erp-service/internal/models"
	"simple-erp-service/internal/repository"
	"strings"
)

// CustomerValidator valida regras de negócio relacionadas a clientes
type CustomerValidator struct {
	customerRepo repository.CustomerRepository
}

// NewCustomerValidator cria um novo validador de clientes
func NewCustomerValidator(customerRepo repository.CustomerRepository) *CustomerValidator {
	return &CustomerValidator{
		customerRepo: customerRepo,
	}
}

// ValidateForCreation valida os dados para criação de um cliente
func (v *CustomerValidator) ValidateForCreation(req models.CreateCustomerRequest) error {
	var errors ValidationErrors

	// Verificar se o documento já existe
	if req.Document != "" {
		// Remover caracteres não numéricos para comparação
		document := strings.ReplaceAll(req.Document, ".", "")
		document = strings.ReplaceAll(document, "-", "")
		document = strings.ReplaceAll(document, "/", "")
		
		exists, err := v.customerRepo.ExistsByDocument(document)
		if err != nil {
			return err
		}
		if exists {
			errors.AddError("document", "documento já está em uso")
		}
	}

	// Verificar se o email já existe (se fornecido)
	if req.Email != "" {
		exists, err := v.customerRepo.ExistsByEmail(req.Email)
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

// ValidateForUpdate valida os dados para atualização de um cliente
func (v *CustomerValidator) ValidateForUpdate(id uint, req models.UpdateCustomerRequest) error {
	var errors ValidationErrors

	// Verificar se o cliente existe
	customer, err := v.customerRepo.FindByID(id)
	if err != nil {
		return err
	}
	if customer == nil {
		errors.AddError("id", "cliente não encontrado")
		return errors
	}

	// Verificar se o documento já está em uso por outro cliente (se fornecido)
	if req.Document != "" && req.Document != customer.Document {
		// Remover caracteres não numéricos para comparação
		document := strings.ReplaceAll(req.Document, ".", "")
		document = strings.ReplaceAll(document, "-", "")
		document = strings.ReplaceAll(document, "/", "")
		
		exists, err := v.customerRepo.ExistsByDocumentExcept(document, id)
		if err != nil {
			return err
		}
		if exists {
			errors.AddError("document", "documento já está em uso")
		}
	}

	// Verificar se o email já está em uso por outro cliente (se fornecido)
	if req.Email != "" && req.Email != customer.Email {
		exists, err := v.customerRepo.ExistsByEmailExcept(req.Email, id)
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