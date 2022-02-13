package converters

import "bitbucket.org/houmeteam/houme-go/models"

func ConvertJobToCompanyJob(job models.Job) models.CompanyJob {
	return models.CompanyJob{
		JobName:                        job.JobName,
		StageName:                      job.StageName,
		SubStageName:                   job.SubStageName,
		WallMaterial:                   job.WallMaterial,
		FinishMaterial:                 job.FinishMaterial,
		FoundationMaterial:             job.FoundationMaterial,
		RoofingMaterial:                job.RoofingMaterial,
		InteriorMaterial:               job.InteriorMaterial,
		ConstructionSpeed:              job.ConstructionSpeed,
		ConstructionCost:               job.ConstructionCost,
		ConstructionFixDurationInHours: job.ConstructionFixDurationInHours,

		JobId: job.JobId,
	}
}
