package repositories

import (
	"log"

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

	log.Println(companyName)

	err := r.db.Where("company_name = ?", companyName).Find(&properties).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &properties}
}
