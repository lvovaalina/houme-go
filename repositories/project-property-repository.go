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
	err := r.db.Unscoped().Where("project_refer = ?", projectId).Delete(&models.ProjectProperty{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ProjectPropertyRepository) FindProjectPropertiesByProjectId(projectId string) RepositoryResult {
	var properties []models.ProjectProperty

	err := r.db.Where("project_refer = ?", projectId).Find(&properties).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &properties}
}
