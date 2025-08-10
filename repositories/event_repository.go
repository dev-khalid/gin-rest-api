package repositories

import (
	"errors"

	"github.com/dev-khalid/gin-rest-api/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// EventRepository defines the contract for event data operations.
type EventRepository interface {
	CreateEvent(event *models.Event) error
	GetEvents() ([]models.Event, error)
	GetEventByID(id uint) (*models.Event, error)
	UpdateEvent(id uint, updated *models.Event) (*models.Event, error)
	DeleteEvent(id uint) error
}

// gormEventRepository implements EventRepository using Gorm.
type gormEventRepository struct {
	db *gorm.DB
}

// NewEventRepository creates a new repository with the provided gorm DB.
func NewEventRepository(db *gorm.DB) EventRepository {
	return &gormEventRepository{db: db}
}

func (r *gormEventRepository) CreateEvent(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *gormEventRepository) GetEvents() ([]models.Event, error) {
	var events []models.Event
	if err := r.db.Preload(clause.Associations).Order("id DESC").Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}

func (r *gormEventRepository) GetEventByID(id uint) (*models.Event, error) {
	var event models.Event
	if err := r.db.Preload(clause.Associations).First(&event, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &event, nil
}

func (r *gormEventRepository) UpdateEvent(id uint, updated *models.Event) (*models.Event, error) {
	// Ensure the record exists
	var existing models.Event
	if err := r.db.First(&existing, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	updated.ID = existing.ID
	// Save will update all fields, including zero-values; use Select if partial updates are needed.
	if err := r.db.Model(&existing).Updates(updated).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

func (r *gormEventRepository) DeleteEvent(id uint) error {
	res := r.db.Delete(&models.Event{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
