package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Repository[models.User]
}

type userRepository struct {
	Repository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		Repository: NewRepository[models.User](db),
		db:         db,
	}
}
