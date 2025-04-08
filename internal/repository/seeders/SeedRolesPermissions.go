package seeders

import (
	"log"

	"simple-erp-service/internal/models"

	"gorm.io/gorm"
)

// SeedRolesPermissions insere perfis e permissões no banco
func SeedRolesPermissions(db *gorm.DB) {
	err := db.Transaction(func(tx *gorm.DB) error {
		// Seed de Roles (Perfis)
		roles := []models.Role{
			{Name: "ADMIN", Description: "Administrador do sistema com acesso completo"},
			{Name: "Gerente", Description: "Acesso gerencial a múltiplos módulos"},
			{Name: "Vendas", Description: "Acesso ao módulo de vendas"},
			{Name: "Estoque", Description: "Acesso ao módulo de estoque"},
			{Name: "Financeiro", Description: "Acesso ao módulo financeiro"},
		}

		for _, role := range roles {
			if err := tx.Where("name = ?", role.Name).FirstOrCreate(&role).Error; err != nil {
				return err
			}
		}

		// Seed de Permissions (Permissões)
		permissions := []models.Permission{
			// Usuários
			{Name: "users.view", Description: "Visualizar usuários", Module: "users"},
			{Name: "users.create", Description: "Criar usuários", Module: "users"},
			{Name: "users.edit", Description: "Editar usuários", Module: "users"},
			{Name: "users.delete", Description: "Excluir usuários", Module: "users"},

			// Vendas
			{Name: "sales.view", Description: "Visualizar vendas", Module: "sales"},
			{Name: "sales.create", Description: "Criar vendas", Module: "sales"},
			{Name: "sales.edit", Description: "Editar vendas", Module: "sales"},
			{Name: "sales.delete", Description: "Excluir vendas", Module: "sales"},
			{Name: "sales.reports", Description: "Gerar relatórios de vendas", Module: "sales"},

			// Estoque
			{Name: "inventory.view", Description: "Visualizar estoque", Module: "inventory"},
			{Name: "inventory.create", Description: "Adicionar itens ao estoque", Module: "inventory"},
			{Name: "inventory.edit", Description: "Editar itens do estoque", Module: "inventory"},
			{Name: "inventory.delete", Description: "Remover itens do estoque", Module: "inventory"},
			{Name: "inventory.reports", Description: "Gerar relatórios de estoque", Module: "inventory"},

			// Financeiro
			{Name: "finance.view", Description: "Visualizar finanças", Module: "finance"},
			{Name: "finance.create", Description: "Criar transações financeiras", Module: "finance"},
			{Name: "finance.edit", Description: "Editar transações financeiras", Module: "finance"},
			{Name: "finance.delete", Description: "Excluir transações financeiras", Module: "finance"},
			{Name: "finance.reports", Description: "Gerar relatórios financeiros", Module: "finance"},
		}

		for _, perm := range permissions {
			if err := tx.Where("name = ?", perm.Name).FirstOrCreate(&perm).Error; err != nil {
				return err
			}
		}

		// Atribuir permissões ao ADMIN
		var adminRole models.Role
		if err := tx.Where("name = ?", "ADMIN").First(&adminRole).Error; err != nil {
			return err
		}

		var allPermissions []models.Permission
		if err := tx.Find(&allPermissions).Error; err != nil {
			return err
		}

		if err := tx.Model(&adminRole).Association("Permissions").Append(&allPermissions); err != nil {
			return err
		}

		// Atribuir permissões específicas por perfil
		assignRolePermissions(tx, "Vendas", "sales")
		assignRolePermissions(tx, "Estoque", "inventory")
		assignRolePermissions(tx, "Financeiro", "finance")

		// Atribuir permissões de visualização e relatórios ao Gerente
		var managerRole models.Role
		if err := tx.Where("name = ?", "Gerente").First(&managerRole).Error; err != nil {
			return err
		}

		var viewPermissions []models.Permission
		if err := tx.Where("name LIKE ?", "%.view").Or("name LIKE ?", "%.reports").Find(&viewPermissions).Error; err != nil {
			return err
		}

		if err := tx.Model(&managerRole).Association("Permissions").Append(&viewPermissions); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatalf("Erro ao executar SeedRolesPermissions: %v", err)
	}

	log.Println("Seeder de Roles e Permissions executado com sucesso!")
}

// assignRolePermissions atribui permissões de um módulo a um perfil específico
func assignRolePermissions(tx *gorm.DB, roleName string, moduleName string) error {
	var role models.Role
	if err := tx.Where("name = ?", roleName).First(&role).Error; err != nil {
		return err
	}

	var modulePermissions []models.Permission
	if err := tx.Where("module = ?", moduleName).Find(&modulePermissions).Error; err != nil {
		return err
	}

	return tx.Model(&role).Association("Permissions").Append(&modulePermissions)
}
