package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type MediaRepository interface {
	Repository[models.Media]
}

type mediaRepository struct {
	Repository[models.Media]
}

func NewMediaRepository(db *gorm.DB) MediaRepository {
	return &mediaRepository{
		Repository: NewRepository[models.Media](db),
	}
}
