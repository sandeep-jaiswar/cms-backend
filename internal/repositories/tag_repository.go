package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type TagRepository interface {
	Repository[models.Tag]
}

type tagRepository struct {
	Repository[models.Tag]
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{
		Repository: NewRepository[models.Tag](db),
	}
}
