package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type CompanyJobRepository struct {
	db *gorm.DB
}

func NewCompanyJobRepository(db *gorm.DB) *CompanyJobRepository {
	return &CompanyJobRepository{db: db}
}

func (r *CompanyJobRepository) Save(companyJob models.CompanyJob) RepositoryResult {
	err := r.db.Save(&companyJob).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &companyJob}
}

func (r *CompanyJobRepository) GetByCompanyId(id string) RepositoryResult {
	var companyJobs []models.CompanyJob

	err := r.db.Find(&companyJobs).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &companyJobs}
}

func (r *CompanyJobRepository) GetById(id string) RepositoryResult {
	var companyJob models.CompanyJob

	err := r.db.Find(&companyJob, id).Error
	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &companyJob}
}

func (r *CompanyJobRepository) Delete(id string) RepositoryResult {
	err := r.db.Delete(&models.CompanyJob{}, id).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
