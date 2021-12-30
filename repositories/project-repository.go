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
	var projects []models.ProjectMin

	err := r.db.Order("project_id").Model(&models.Project{}).Find(&projects).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &projects}
}

func (r *ProjectRepository) FindAllProjects() RepositoryResult {
	var projects []models.Project

	err := r.db.Preload("ProjectJobs").
		Preload("ProjectJobs.Job").
		Preload("ProjectJobs.Job.Property").
		Preload("ProjectProperties").
		Preload("ProjectProperties.Property").
		Find(&projects).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &projects}
}

func (r *ProjectRepository) DeleteProjectById(id string) RepositoryResult {
	err := r.db.Unscoped().Delete(&models.Project{}, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ProjectRepository) GetProjectById(id string) RepositoryResult {
	var project models.Project
	err := r.db.Preload("ProjectJobs").
		Preload("ProjectJobs.Job").
		Preload("ProjectJobs.Job.Property").
		Preload("ProjectProperties").
		Preload("ProjectProperties.Property").
		Preload("ProjectMaterials").
		Preload("ProjectMaterials.ConstructionJobMaterial").
		Preload("ProjectMaterials.ConstructionJobMaterial.Job").
		Preload("ProjectMaterials.ConstructionJobMaterial.Job.Property").
		First(&project, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &project}
}
