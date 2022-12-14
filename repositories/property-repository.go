package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type PropertyRepository struct {
	db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) *PropertyRepository {
	return &PropertyRepository{db: db}
}

func (r *PropertyRepository) FindAll() RepositoryResult {
	var properties []models.Property

	err := r.db.Find(&properties).Order("id").Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &properties}
}
