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

	// Injeção dos services
	userService := service.GetInstanceUser()
	invoiceService := service.GetInstanceInvoice()
	clientService := service.GetInstanceClient()
	signatureService := service.GetInstanceSignature()
	operacaoService := service.GetInstanceOperacao()
	usuarioService := service.GetInstanceUsuario()
	expenseCenterService := service.GetInstanceExpenseCenter()

	// Injeção dos handlers
	userHandler := handler.NewUserHandler(userService)
	invoiceHandler := handler.NewInvoiceHandler(invoiceService)
	clientHandler := handler.NewClientHandler(clientService, userService)
	signatureHandler := handler.NewSignatureHandler(signatureService, clientService, expenseCenterService)
	operacaoHandler := handler.NewOperacaoHandler(operacaoService)
	usuarioHandler := handler.NewUsuarioHandler(usuarioService)
	expenseCenterHandler := handler.NewExpenseCenterHandler(expenseCenterService, userService)

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

		protected.POST("/user", userHandler.GetUserByID)
		protected.PUT("/user", userHandler.UpdateUser)
		protected.DELETE("/user", userHandler.DeleteUser)
		protected.GET("/users", userHandler.GetAllUsers)

		// Client
		protected.POST("/client", clientHandler.CreateClient)
		protected.GET("/client/id", clientHandler.GetClientByID)
		protected.GET("/client/user", clientHandler.GetClientByUserID)
		protected.GET("/client", clientHandler.GetAllClient)
		protected.PUT("/client/id/:id", clientHandler.UpdateClient)
		protected.DELETE("/client", clientHandler.DeleteClient)
		protected.POST("/client/cpf", clientHandler.GetClientByCpf)
		protected.POST("/client/UF", clientHandler.GetClientByUF)
		protected.POST("/client/cidade", clientHandler.GetClientByCidade)

		// Signature
		protected.POST("/signature", signatureHandler.CreateSignature)
		protected.GET("/signature/id/:id", signatureHandler.GetSignatureByID)
		protected.GET("/signature", signatureHandler.GetAllSignature)
		protected.PUT("/signature", signatureHandler.UpdateSignature)
		protected.DELETE("/signature", signatureHandler.DeleteSignature)
		protected.POST("/signature/cliente", signatureHandler.GetSignatureByClienteID)

		// Operacao
		protected.POST("/operacao", operacaoHandler.CreateOperacao)
		protected.GET("/operacao/id/:id", operacaoHandler.GetOperacaoByID)
		protected.GET("/operacao", operacaoHandler.GetAllOperacao)
		protected.PUT("/operacao/id/:id", operacaoHandler.UpdateOperacao)
		protected.DELETE("/operacao", operacaoHandler.DeleteOperacao)

		// Usuario
		protected.POST("/usuario", usuarioHandler.CreateUsuario)
		protected.GET("/usuario/id/:id", usuarioHandler.GetUsuarioByID)
		protected.GET("/usuario", usuarioHandler.GetAllUsuario)
		protected.PUT("/usuario/id/:id", usuarioHandler.UpdateUsuario)
		protected.DELETE("/usuario", usuarioHandler.DeleteUsuario)
		protected.POST("/usuario/email", usuarioHandler.GetUsuarioByEmail)

		protected.POST("/expensecenter", expenseCenterHandler.CreateExpenseCenter)
		protected.GET("/expensecenter/id/:id", expenseCenterHandler.GetExpenseCenterByID)
		protected.GET("/expensecenter/user", expenseCenterHandler.GetExpenseCenterByUserID)
		protected.PUT("/expensecenter", expenseCenterHandler.UpdateExpenseCenter)
		protected.DELETE("/expensecenter", expenseCenterHandler.DeleteExpenseCenter)
		protected.GET("/expensecenter", expenseCenterHandler.GetAllExpenseCenter)

	}

	// Inicia servidor na porta 9090
	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
