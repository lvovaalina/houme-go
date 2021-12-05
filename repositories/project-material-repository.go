package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type ProjectMaterialRepository struct {
	db *gorm.DB
}

func NewProjectMaterialRepository(db *gorm.DB) *ProjectMaterialRepository {
	return &ProjectMaterialRepository{db: db}
}

func (r *ProjectMaterialRepository) DeleteProjectMaterialsByProjectId(projectId string) RepositoryResult {
	err := r.db.Unscoped().Where("project_refer = ?", projectId).Delete(&models.ProjectJobMaterial{}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
