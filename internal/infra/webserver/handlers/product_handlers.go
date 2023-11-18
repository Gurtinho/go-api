package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gurtinho/go/api/internal/dtos"
	"github.com/gurtinho/go/api/internal/entities"
	"github.com/gurtinho/go/api/internal/infra/database"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (prod *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dtos.ProductDTO
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entities.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = prod.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}