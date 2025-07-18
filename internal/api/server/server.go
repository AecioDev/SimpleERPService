package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"simple-erp-service/config"
	"simple-erp-service/internal/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Server representa o servidor HTTP
type Server struct {
	router *gin.Engine
	cfg    *config.Config
	db     *gorm.DB
}

// NewServer cria uma nova instância do servidor
func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	router := gin.Default()

	// Configurar CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000", "http://localhost:3001"}, // <--- MUDANÇA AQUI: Especifique a origem do seu frontend
		// Se você tiver múltiplos frontends ou um domínio de produção:
		// AllowOrigins: []string{"http://localhost:3000", "https://seu-dominio-frontend.com"},

		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // <--- Isso deve permanecer TRUE
		MaxAge:           12 * time.Hour,
	}))

	return &Server{
		router: router,
		cfg:    cfg,
		db:     db,
	}
}

// Run inicia o servidor HTTP
func (s *Server) Run() error {
	// Configurar Swagger
	s.setupSwagger()

	// Configurar rotas
	s.setupRoutes()

	// Configurar servidor HTTP
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", s.cfg.Server.Port),
		Handler:      s.router,
		ReadTimeout:  s.cfg.Server.ReadTimeout,
		WriteTimeout: s.cfg.Server.WriteTimeout,
		IdleTimeout:  s.cfg.Server.IdleTimeout,
	}

	// Canal para capturar sinais de interrupção
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Iniciar servidor em uma goroutine
	go func() {
		log.Printf("Servidor iniciado na porta %s", s.cfg.Server.Port)
		log.Printf("Documentação Swagger disponível em http://localhost:%s/swagger/index.html", s.cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar servidor: %v", err)
		}
	}()

	// Aguardar sinal de interrupção
	<-quit
	log.Println("Desligando servidor...")

	// Contexto com timeout para desligamento gracioso
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Desligar servidor
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Erro ao desligar servidor: %v", err)
	}

	log.Println("Servidor desligado com sucesso")
	return nil
}

// setupRoutes configura todas as rotas da API
func (s *Server) setupRoutes() {
	// Grupo de rotas da API
	api := s.router.Group("/api")

	// Configurar rotas para cada módulo
	routes.SetupAuthRoutes(api, s.db, s.cfg)
	routes.SetupUserRoutes(api, s.db)
	routes.SetupRoleRoutes(api, s.db)
	routes.SetupProductsRoutes(api, s.db)
	routes.SetupInventoryRoutes(api, s.db)
	routes.SetupCustomersRoutes(api, s.db)
	routes.SetupSupplierRoutes(api, s.db)
	routes.SetupSalesRoutes(api, s.db)
	routes.SetupPurchasesRoutes(api, s.db)
	routes.SetupFinancialRoutes(api, s.db)
	routes.SetupDashboardRoutes(api, s.db)
	routes.SetupSystemRoutes(api, s.db)
	routes.SetupPermissionRoutes(api, s.db)
}
