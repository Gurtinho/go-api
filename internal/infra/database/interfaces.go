package database

import (
	"github.com/gurtinho/go/api/internal/entities"
)

type UserInterface interface {
	create(user *entities.User) (*entities.User, error)
	findByEmail(email string) (*entities.User, error)
}