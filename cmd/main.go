package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()

	if err != nil {
		panic(err)
	}
	productUseCase := usecase.NewProductUseCase()

	ProductController := controller.NewProductController(productUseCase)

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
