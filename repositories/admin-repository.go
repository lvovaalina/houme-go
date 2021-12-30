package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) Save(admin models.Admin) RepositoryResult {
	err := r.db.Save(&admin).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &admin}
}

func (r *AdminRepository) GetByEmail(email string) RepositoryResult {
	var admin models.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &admin}
}
