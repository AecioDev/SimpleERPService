package db

import (
	"simple-erp-service/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB inicializa a conexão com o banco de dados
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.Database.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Registrar modelos para auto-migração se necessário
	// Nota: Como já temos o esquema SQL, não precisamos de auto-migração
	// mas é bom ter isso configurado para desenvolvimento

	return db, nil
}
