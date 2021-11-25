package converters

import "bitbucket.org/houmeteam/houme-go/models"

func ConvertJobToProjectJob(job models.Job) models.ProjectJob {
	return models.ProjectJob{
		JobId: job.JobId,
		Job:   job,
	}
}
