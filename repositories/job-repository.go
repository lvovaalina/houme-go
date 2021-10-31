package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"github.com/jinzhu/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) FindAll() RepositoryResult {
	var jobs []models.Job

	err := r.db.Find(&jobs).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &jobs}
}
