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
		if jobProp.ConstructionFixDurationInHours > 0.0 {
			constrDur = jobProp.ConstructionFixDurationInHours
		} else {
			constrDur = calculateDuration(
				jobValue, jobProp.ConstructionSpeed, constrWorkNumber)
		}

		constrDurInHours := int(math.Round(float64(constrDur)))
		constrDurInDays := int(math.Round(float64(constrDurInHours) / workingHoursInDay))
		if (constrDurInHours % workingHoursInDay) > 0 {
			constrDurInDays += 1
		}

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

func CalculateProjectDuration(projectJobs []models.ProjectJob) int {
	var projectDuration int
	for _, j := range projectJobs {
		if !j.Job.InParallel {
			projectDuration += j.ConstructionDurationInDays
		}
	}

	log.Println("Project duration without par: ", projectDuration)

	parGroupMap := map[string][]int{}
	for _, job := range projectJobs {
		if job.Job.InParallel {
			if _, ok := parGroupMap[job.Job.ParallelGroupCode]; ok {
				parGroupMap[job.Job.ParallelGroupCode] = append(parGroupMap[job.Job.ParallelGroupCode], job.ConstructionDurationInDays)
			} else {
				parGroupMap[job.Job.ParallelGroupCode] = []int{job.ConstructionDurationInDays}
			}
		}
	}

	log.Println(parGroupMap)

	for _, group := range parGroupMap {
		projectDuration += findMax(group)
	}

	return projectDuration
}

func findMax(a []int) int {
	max := a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
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
