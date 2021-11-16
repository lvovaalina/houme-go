package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type ProjectPropertyRepository struct {
	db *gorm.DB
}

func NewProjectPropertyRepository(db *gorm.DB) *ProjectPropertyRepository {
	return &ProjectPropertyRepository{db: db}
}

func (r *ProjectPropertyRepository) DeleteProjectPropertiesByProjectId(projectId string) RepositoryResult {
	err := r.db.Where("project_refer = ?", projectId).Delete(&models.ProjectProperty{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
