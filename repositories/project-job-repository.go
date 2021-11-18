package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type ProjectJobRepository struct {
	db *gorm.DB
}

func NewProjectJobRepository(db *gorm.DB) *ProjectJobRepository {
	return &ProjectJobRepository{db: db}
}

func (r *ProjectJobRepository) DeleteProjectJobsByProjectId(projectId int) RepositoryResult {

	err := r.db.Where("project_refer = ?", projectId).Delete(&models.ProjectJob{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ProjectJobRepository) FindProjectJobsByProjectId(projectId string) RepositoryResult {
	var jobs []models.ProjectJob

	err := r.db.Where("project_refer = ?", projectId).Preload("Job").Preload("Job.Property").Find(&jobs).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &jobs}
}
