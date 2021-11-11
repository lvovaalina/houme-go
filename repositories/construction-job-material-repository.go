package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type ConstructionJobMaterialRepository struct {
	db *gorm.DB
}

func NewConstructionJobMaterialRepository(db *gorm.DB) *ConstructionJobMaterialRepository {
	return &ConstructionJobMaterialRepository{db: db}
}

func (r *ConstructionJobMaterialRepository) FindMaterials() RepositoryResult {
	var materials []models.ConstructionJobMaterial

	err := r.db.Preload("Job").Preload("Property").Find(&materials).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &materials}
}

func (r *ConstructionJobMaterialRepository) FindJobMaterialById(id string) RepositoryResult {
	var material models.ConstructionJobMaterial

	err := r.db.Find(&material, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &material}
}

func (r *ConstructionJobMaterialRepository) DeleteJobMaterialById(id string) RepositoryResult {
	err := r.db.Delete(&models.ConstructionJobMaterial{}, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}

func (r *ConstructionJobMaterialRepository) Save(jobMaterial *models.ConstructionJobMaterial) RepositoryResult {
	err := r.db.Save(jobMaterial).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: jobMaterial}
}
