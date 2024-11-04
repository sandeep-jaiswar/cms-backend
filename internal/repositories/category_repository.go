package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
    Repository[models.Category]
}

type categoryRepository struct {
    Repository[models.Category]
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
    return &categoryRepository{
        Repository: NewRepository[models.Category](db),
    }
}
