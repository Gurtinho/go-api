package database

import (
	"github.com/gurtinho/go/api/internal/entities"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}

func (usr *User) Create(user *entities.User) error {
	return usr.DB.Create(user).Error
}

func (usr *User) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := usr.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

