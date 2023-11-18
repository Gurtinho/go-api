package database

import (
	"github.com/gurtinho/go/api/internal/entities"
)

type UserInterface interface {
	Create(user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
}

type ProductInterface interface {
	Create(product *entities.Product) error
	FindAll(page, limit int, sort string) ([]entities.Product, error)
	FindByID(ID string) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(ID string) error
}