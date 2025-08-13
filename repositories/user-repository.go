package repositories

import (
	"github.com/dev-khalid/gin-rest-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetAll() ([]*models.User, error)
	// Update(user *models.User) error
	Delete(id uint) error
}

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepository{
		db: db,
	}
}

func (r *gormUserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *gormUserRepository) GetByID(id uint) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *gormUserRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *gormUserRepository) Delete(id uint) error {
	r.db.Scopes()
	return r.db.Select(clause.Associations).Delete(&models.User{}, id).Error
}