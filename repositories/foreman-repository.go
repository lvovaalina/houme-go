package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type ForemanRepository struct {
	db *gorm.DB
}

func NewForemanRepository(db *gorm.DB) *ForemanRepository {
	return &ForemanRepository{db: db}
}

func (r *ForemanRepository) Save(foreman models.Foreman) RepositoryResult {
	err := r.db.Save(&foreman).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &foreman}
}

func (r *ForemanRepository) GetByEmail(email string) RepositoryResult {
	var foreman models.Foreman
	err := r.db.Where("email = ?", email).First(&foreman).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &foreman}
}
