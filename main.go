package main

import (
	"context"
	"fmt"
	"log"
	"pix-generation/src/client"
	"pix-generation/src/handler"
	"pix-generation/src/middleware"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Errorf("Error loading .env file")
	}

	defer cancel()

	if err := client.GetInstance().Initialize(ctx); err != nil {
		fmt.Errorf("mongo off")
	}

	r := gin.Default()

	r.POST("/login", handler.ValidateUser)
	r.POST("/register", handler.CreateUser)

	protected := r.Group("/", middleware.JWTMiddleware())
	{
		protected.POST("/invoice", handler.CreateInvoice)
		protected.GET("/invoice/id/:id", handler.GetByID)
		protected.POST("/invoice/:startDate/:endDate/", handler.GetByCnpj)
		protected.POST("/invoice/cnpj", handler.GetByCnpj)
		protected.DELETE("/invoice/:startDate/:endDate/", handler.DeleteInvoice)

	}

	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
