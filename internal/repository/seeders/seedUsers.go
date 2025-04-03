package seeders

import (
	"log"

	"simple-erp-service/internal/models"

	"gorm.io/gorm"

	"simple-erp-service/internal/utils"
)

func SeedUserAdm(db *gorm.DB) {
	// Hash da senha
	passwordHash, err := utils.HashPassword("987321")
	if err != nil {
		log.Fatal("Erro ao configurar a senha de adm:", err)
	}

	// Buscar o perfil ADMIN
	var adminRole models.Role
	if err := db.Where("name = ?", "ADMIN").First(&adminRole).Error; err != nil {
		log.Fatal("Erro ao recuperar o perfil ADMIN:", err)
	}

	// Criar usuário administrador
	adminUser := models.User{
		Username:     "admin",
		PasswordHash: passwordHash,
		Name:         "Administrador",
		Email:        "admin@sistema.com",
		RoleID:       adminRole.ID,
		IsActive:     true,
	}

	// Evitar duplicação
	if err := db.Where("username = ?", adminUser.Username).FirstOrCreate(&adminUser).Error; err != nil {
		log.Printf("Erro ao criar usuário ADMIN: %v", err)
	} else {
		log.Println("Usuário ADMIN criado ou já existente!")
	}
}
