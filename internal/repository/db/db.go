package db

import (
	"fmt"
	"log"
	"simple-erp-service/config"
	"simple-erp-service/internal/repository/seeders"
	"simple-erp-service/migrations"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB inicializa a conexão com o banco de dados
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}

	// Executar migrações
	if err := migrations.MigrateDB(db); err != nil {
		return nil, err
	}

	// Perguntar ao usuário se deseja rodar os seeders
	var resposta string
	fmt.Print("Deseja rodar os seeders? (s/n): ")
	fmt.Scanln(&resposta)

	if strings.ToLower(strings.TrimSpace(resposta)) == "s" {
		log.Println("Rodando seeders...")
		seeders.SeedRolesPermissions(db)
		log.Println("Seeders concluídos com sucesso!")
	} else {
		log.Println("Seeders ignorados.")
	}

	return db, nil
}
