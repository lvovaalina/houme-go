package helpers

import (
	"log"
	"math"

	"bitbucket.org/houmeteam/houme-go/models"
)

const workingHoursInDay = 8

type JobCalculations struct {
	ConstructionWorkers int
	ConstructionCost    float32

	ConstructionDuration        float32
	ConstructionDurationInHours int
	ConstructionDurationInDays  int

	JobCode string
}

func CalculateCostDurationForProjectJobs(
	project models.Project,
	constructionProperties *[]models.ConstructionJobProperty) []JobCalculations {

	jobPropMap := map[string]models.ConstructionJobProperty{}
	for _, prop := range *constructionProperties {
		jobPropMap[prop.JobCode] = prop
	}

	propertiesMap := map[string]float32{}
	for _, p := range project.ProjectProperties {
		propertiesMap[p.PropertyCode] = p.PropertyValue
	}

	var calcProjectJobs []JobCalculations
	for _, job := range project.ProjectJobs {
		jobProp := jobPropMap[job.JobCode]
		jobValue := propertiesMap[job.PropertyCode]
		constrWorkNumber := calculateWorkers(jobProp, project.ConstructionWorkersNumber)

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

		calcJob := JobCalculations{
			ConstructionWorkers:         constrWorkNumber,
			ConstructionDuration:        constrDur,
			ConstructionDurationInHours: constrDurInHours,
			ConstructionDurationInDays:  constrDurInDays,
			ConstructionCost:            0,
			JobCode:                     job.JobCode,
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
