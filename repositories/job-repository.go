package repositories

import (
	"bitbucket.org/houmeteam/houme-go/models"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db: db}
}

func (r *JobRepository) FindAll() RepositoryResult {
	var jobs []models.Job

	err := r.db.Preload("Property").Find(&jobs).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &jobs}
}

func (r *JobRepository) FindProjectJobs(
	wallMaterial string,
	foundationMaterial string,
	roofingMaterial string,
	finishMaterial string) RepositoryResult {
	var jobs []models.Job

	err := r.db.
		Preload("Property").
		Where(
			r.db.Where("required = ?", false).Where(
				r.db.Where("foundation_material = ?", foundationMaterial).
					Or("wall_material = ?", wallMaterial).
					Or("roofing_material = ?", roofingMaterial).
					Or("finish_material = ?", finishMaterial))).
		Or("required = ?", true).
		Find(&jobs).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &jobs}
}
