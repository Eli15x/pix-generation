package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"pix-generation/src/client"
	"pix-generation/src/handler"
	"pix-generation/src/middleware"


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

	protected := r.Group("/", middleware.JWTMiddleware())
	{
		protected.POST("/invoice", handler.UpdateInvoice)
		protected.GET("/invoice/:id", handler.GetInvoice)
		protected.GET("/invoice/:startDate/:endDate/", handler.GetInvoice)
		protected.DELETE("/invoice", handler.DeleteInvoice)

	}

	if err := r.Run(":9090"); err != nil {
		log.Fatal(err)
	}
}
