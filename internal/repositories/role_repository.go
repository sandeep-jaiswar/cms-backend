package repositories

import (
	"github.com/sandeep-jaiswar/cms-backend/internal/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Repository[models.Role]
}

type roleRepository struct {
	Repository[models.Role]
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &roleRepository{
		Repository: NewRepository[models.Role](db),
	}
}
