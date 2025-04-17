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

	// Rotas públicas
	r.POST("/login", handler.ValidateUser)
	r.POST("/register", handler.CreateUser)

	// Rotas protegidas com JWT
	protected := r.Group("/", middleware.JWTMiddleware())
	{
		protected.POST("/invoice", handler.CreateInvoice)
		protected.GET("/invoice/id/:id", handler.GetByID)
		protected.POST("/invoice/:startDate/:endDate/", handler.GetByCnpj)
		protected.POST("/invoice/cnpj", handler.GetByCnpj)
		protected.DELETE("/invoice/:startDate/:endDate/", handler.DeleteInvoice)
		protected.GET("/user", handler.GetUserByID)
		protected.PUT("/user", handler.UpdateUser)
		protected.DELETE("/user", handler.DeleteUser)
		protected.GET("/users", handler.GetAllUsers)
	}

	// Inicia servidor na porta 9090
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
