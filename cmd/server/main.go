package main

import (
	"net/http"

	"github.com/gurtinho/go/api/configs"
	"github.com/gurtinho/go/api/internal/entities"
	"github.com/gurtinho/go/api/internal/infra/database"
	"github.com/gurtinho/go/api/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig("./cmd/server")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("./cmd/server/test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entities.Product{}, &entities.User{})

  // puxando o database
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	// handlers
	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":8000", nil)
}