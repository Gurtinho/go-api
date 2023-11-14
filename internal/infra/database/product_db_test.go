package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/gurtinho/go/api/internal/entities"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("product1", 100.0)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entities.Product{})
	for i := 1; i < 24; i++ {
		product, err := entities.NewProduct(fmt.Sprintf("product %d", i), rand.Float64() * 100)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)

	// 1ª Page
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "product 1", products[0].Name)
	assert.Equal(t, "product 10", products[9].Name)

	// 2ª page
	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "product 11", products[0].Name)
	assert.Equal(t, "product 20", products[9].Name)

	// 3ª page
	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 3)
	assert.Equal(t, "product 21", products[0].Name)
	assert.Equal(t, "product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("camiseta", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "camiseta", product.Name)
	assert.Equal(t, 10.0, product.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("camiseta", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	product.Name = "camiseta 2"
	err = productDB.Update(product)
	assert.NoError(t, err)
	product, err = productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "camiseta 2", product.Name)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entities.Product{})
	product, err := entities.NewProduct("camiseta", 10.0)
	assert.NoError(t, err)
	db.Create(product)
	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)
	_, err = productDB.FindByID(product.ID.String())
	assert.Error(t, err)
}