package routes

import (
	"simple-erp-service/config"
	"simple-erp-service/internal/api/handlers"
	"simple-erp-service/internal/api/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupFinancialRoutes configura as rotas de financial
func SetupFinancialRoutes(router *gin.RouterGroup, db *gorm.DB) {
	financialHandler := handlers.NewFinancialHandler(db)

	// Obter configuração para middleware de autenticação
	cfg, _ := config.Load()

	// Grupo de rotas de usuários (todas protegidas)
	transactions := router.Group("/transactions")
	transactions.Use(middlewares.AuthMiddleware(cfg))
	{
		transactions.GET("", middlewares.RequirePermission("financial.view"), financialHandler.GetTransactions)
		transactions.GET("/:id", middlewares.RequirePermission("transactions.view"), financialHandler.GetTransaction)
		transactions.POST("", middlewares.RequirePermission("transactions.create"), financialHandler.CreateTransaction)
		transactions.PUT("/:id", middlewares.RequirePermission("transactions.edit"), financialHandler.UpdateTransaction)
		transactions.DELETE("/:id", middlewares.RequirePermission("transactions.delete"), financialHandler.DeleteTransaction)
		transactions.PUT("/:id/password", financialHandler.ChangePassword) // Permissão verificada no handler
	}
}
