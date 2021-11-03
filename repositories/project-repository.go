package repositories

import (
	"gorm.io/gorm"

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

func (r *ProjectRepository) FindAll() RepositoryResult {
	var properties []models.ProjectMin

	err := r.db.Model(&models.Project{}).Find(&properties).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &properties}
}

func (r *ProjectRepository) DeleteProjectById(id string) RepositoryResult {
	err := r.db.Delete(&models.Project{}, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ProjectRepository) GetProjectById(id string) RepositoryResult {
	var project models.Project
	err := r.db.Preload("ProjectJobs").
		Preload("ProjectJobs.Job").
		Preload("ProjectProperties").
		Preload("ProjectProperties.Property").
		First(&project, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &project}
}
