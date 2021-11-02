package helpers

import (
	"log"
	"math"

	"bitbucket.org/houmeteam/houme-go/models"
)

const workingHoursInDay = 8

func CalculateCostDurationForProjectJobs(
	project models.Project,
	constructionProperties *[]models.ConstructionJobProperty) []models.ProjectJob {

	jobPropMap := map[string]models.ConstructionJobProperty{}
	for _, prop := range *constructionProperties {
		jobPropMap[prop.JobCode] = prop
	}

	propertiesMap := map[string]float32{}
	for _, p := range project.ProjectProperties {
		propertiesMap[p.PropertyCode] = p.PropertyValue
	}

	var calcProjectJobs []models.ProjectJob
	for _, job := range project.ProjectJobs {
		jobProp := jobPropMap[job.JobCode]
		jobValue := propertiesMap[job.PropertyCode]
		constrWorkNumber := calculateWorkers(jobProp, project.ConstructonWorkersNumber)

		var constrDur float32

		log.Println("Constr Fix Dur", jobProp.ConstructionFixDurationInHours)
		if jobProp.ConstructionFixDurationInHours > 0.0 {
			constrDur = jobProp.ConstructionFixDurationInHours
			log.Println("Constr Fix Dur", jobProp.ConstructionFixDurationInHours)
		} else {
			constrDur = calculateDuration(
				jobValue, jobProp.ConstructionSpeed, constrWorkNumber)
		}

		constrDurInHours := int(math.Round(float64(constrDur)))
		constrDurInDays := int(math.Round(float64(constrDurInHours) / workingHoursInDay))

		log.Println("Job Code: ", job.JobCode,
			",Work Num: ", constrWorkNumber, ", Dur: ", constrDur)

		calcJob := models.ProjectJob{
			ConstructionWorkers:     constrWorkNumber,
			ConstructionDuration:    constrDur,
			ConstructionCostInHours: constrDurInHours,
			ConstructionCostInDays:  constrDurInDays,
			ConstructionCost:        0,
			JobId:                   job.JobId,
		}

		calcProjectJobs = append(calcProjectJobs, calcJob)
	}
	return calcProjectJobs
}

func calculateWorkers(jobProperties models.ConstructionJobProperty, constructionWorkersNumber string) int {
	switch constructionWorkersNumber {
	case "Maximum":
		return jobProperties.MaxWorkers
	case "Optimal":
		return jobProperties.OptWorkers
	case "Minimum":
		return jobProperties.MinWorkers
	default:
		return jobProperties.OptWorkers
	}
}

func calculateDuration(value float32, speed float32, numberOfWorkers int) float32 {
	return value / (speed * float32(numberOfWorkers))
}

func calculateCost(constructionDurationInHours int, constructionCost float32) float32 {
	return float32(constructionDurationInHours) * constructionCost
}
