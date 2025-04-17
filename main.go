package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// importa o pacote para ativar o init()
	_ "pix-generation/docs"
	"pix-generation/src/client"
	"pix-generation/src/handler"
	"pix-generation/src/middleware"
	"pix-generation/src/service"

	"pix-generation/src/metrics"
)

// @title           Pix Generation API
// @version         1.0
// @description     API para controle de usuários e invoices com autenticação JWT.
// @termsOfService  http://swagger.io/terms/
// @license.name    Konachse
// @host            localhost:9090
// @BasePath        /
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Carrega variáveis de ambiente
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Aviso: .env não carregado ou inexistente")
	}

	// Inicializa MongoDB
	if err := client.GetInstance().Initialize(ctx); err != nil {
		log.Fatalf("Erro ao conectar no MongoDB: %v", err)
	}

	// Inicia o Gin
	r := gin.Default()

	//Prometheus
	metrics.Init()
	r.Use(metrics.PrometheusMiddleware())
	r.GET("/metrics", metrics.PrometheusHandler())

	// Rota da documentação Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userService := service.GetInstanceUser()
	invoiceService := service.GetInstanceInvoice()

	// Injeção dos handlers
	userHandler := handler.NewUserHandler(userService)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)

	// Rotas públicas
	r.POST("/login", userHandler.ValidateUser)
	r.POST("/register", userHandler.CreateUser)

	// Rotas protegidas com JWT
	protected := r.Group("/", middleware.JWTMiddleware())
	{
		protected.POST("/invoice", invoiceHandler.CreateInvoice)
		protected.GET("/invoice/id/:id", invoiceHandler.GetByID)
		protected.POST("/invoice/:startDate/:endDate/", invoiceHandler.GetByCnpj)
		protected.POST("/invoice/cnpj", invoiceHandler.GetByCnpj)
		protected.DELETE("/invoice/:startDate/:endDate/", invoiceHandler.DeleteInvoice)
		protected.GET("/user", userHandler.GetUserByID)
		protected.PUT("/user", userHandler.UpdateUser)
		protected.DELETE("/user", userHandler.DeleteUser)
		protected.GET("/users", userHandler.GetAllUsers)
	}

	// Inicia servidor na porta 9090
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
