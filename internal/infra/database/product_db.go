package database

import (
	"strings"

	"github.com/gurtinho/go/api/internal/entities"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (prod *Product) Create(product *entities.Product) error {
	return prod.DB.Create(product).Error
}

func (prod *Product) FindAll(page, limit int, sort string) ([]entities.Product, error) {
	var products []entities.Product
	var err error
	sort = strings.ToLower(sort)
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = prod.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = prod.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (prod *Product) FindByID(ID string) (*entities.Product, error) {
	var product entities.Product
	err := prod.DB.First(&product, "id = ?", ID).Error
	return &product, err
}

func (prod *Product) Update(product *entities.Product) error {
	_, err := prod.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return prod.DB.Save(product).Error
}

func (prod *Product) Delete(ID string) error {
	product, err := prod.FindByID(ID)
	if err != nil {
		return err
	}
	return prod.DB.Delete(product).Error
}