package migrations

import (
	"log"
	"simple-erp-service/internal/models"

	"gorm.io/gorm"
)

// MigrateDB executa as migrações do banco de dados
func MigrateDB(db *gorm.DB) error {
	log.Println("Iniciando migrações do banco de dados...")

	// Lista de todos os modelos para migração
	models := []interface{}{
		&models.Account{},
		&models.Address{},
		&models.City{},
		&models.Contact{},
		&models.Country{},
		&models.Customer{},
		&models.Document{},
		&models.FinancialTransaction{},
		&models.InventoryMovement{},
		&models.MeasurementUnit{},
		&models.PaymentMethod{},
		&models.Payment{},
		&models.Permission{},
		&models.ProductCategory{},
		&models.Product{},
		&models.PurchaseItem{},
		&models.Purchase{},
		&models.Role{},
		&models.SaleItem{},
		&models.Sale{},
		&models.State{},
		&models.Supplier{},
		&models.SystemLog{},
		&models.User{},
		&models.Transaction{},
	}

	// Executar migrações
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Printf("Erro ao executar migrações: %v", err)
		return err
	}

	log.Println("Migrações concluídas com sucesso!")
	return nil
}
