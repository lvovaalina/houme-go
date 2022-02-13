package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (r *CompanyRepository) Save(company models.Company) RepositoryResult {
	err := r.db.Save(&company).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &company}
}

func (r *CompanyRepository) Get() RepositoryResult {
	var companies []models.Company

	err := r.db.Preload("Foremen").Find(&companies).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &companies}
}

func (r *CompanyRepository) Delete(id string) RepositoryResult {
	err := r.db.Delete(&models.Company{}, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
