package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type ConstructionJobPropertyRepository struct {
	db *gorm.DB
}

func NewConstructionJobPropertyRepository(db *gorm.DB) *ConstructionJobPropertyRepository {
	return &ConstructionJobPropertyRepository{db: db}
}

func (r *ConstructionJobPropertyRepository) FindPropertiesByCompanyName(companyName string) RepositoryResult {
	var properties []models.ConstructionJobProperty

	err := r.db.Where("company_name = ?", companyName).Preload("Job").Find(&properties).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &properties}
}

func (r *ConstructionJobPropertyRepository) FindJobPropertyById(id string) RepositoryResult {
	var property models.ConstructionJobProperty

	err := r.db.Find(&property, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &property}
}

func (r *ConstructionJobPropertyRepository) Save(jobProperty *models.ConstructionJobProperty) RepositoryResult {
	err := r.db.Save(jobProperty).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: jobProperty}
}
