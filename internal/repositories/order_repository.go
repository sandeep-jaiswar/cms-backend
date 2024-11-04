package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Repository[models.Order]
}

type orderRepository struct {
	Repository[models.Order]
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		Repository: NewRepository[models.Order](db),
		db:         db,
	}
}
