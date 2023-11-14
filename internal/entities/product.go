package entities

import (
	"errors"
	"time"
	"github.com/gurtinho/go/api/pkg/entity"
)

var (
	ErrInvalidID = errors.New("invalid ID")
	ErrInvalidPrice = errors.New("invalid price")
	ErrIDIsRequired = errors.New("id is required")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
}

func (product *Product) ValidadeProduct() error {
	if product.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(product.ID.String()); err != nil {
		return ErrInvalidID
	}
	if product.Name == "" {
		return ErrNameIsRequired
	}
	if product.Price == 0 {
		return ErrPriceIsRequired
	}
	if product.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	err := product.ValidadeProduct()
	if err != nil {
		return nil, err
	}
	return product, nil
}
