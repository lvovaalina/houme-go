package repositories

import (
	"github.com/jinzhu/gorm"

	"bitbucket.org/houmeteam/houme-go/models"
)

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) Save(project *models.Project) RepositoryResult {
	err := r.db.Save(project).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: project}
}
