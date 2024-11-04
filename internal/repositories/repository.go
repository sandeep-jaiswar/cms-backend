package repositories

import "gorm.io/gorm"

type Repository[T any] interface {
	Create(entity *T) error
	FindByID(id uint) (*T, error)
	FindAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
}

type repository[T any] struct {
	db *gorm.DB
}

var (
	ProductRepo  ProductRepository
	UserRepo     UserRepository
	CategoryRepo CategoryRepository
	OrderRepo    OrderRepository
)

func InitRepositories(db *gorm.DB) {
	ProductRepo = NewProductRepository(db)
	UserRepo = NewUserRepository(db)
	CategoryRepo = NewCategoryRepository(db)
	OrderRepo = NewOrderRepository(db)
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{db}
}

func (r *repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *repository[T]) FindByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (r *repository[T]) FindAll() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *repository[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}
