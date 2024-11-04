package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Repository[models.Product]
}

type productRepository struct {
	Repository[models.Product]
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{
		Repository: NewRepository[models.Product](db),
		db:         db,
	}
}
