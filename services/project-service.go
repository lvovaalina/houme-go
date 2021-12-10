package services

import (
	"encoding/base64"
	"log"
	"strconv"

	"bitbucket.org/houmeteam/houme-go/converters"
	"bitbucket.org/houmeteam/houme-go/dtos"
	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
)

func CreateProject(
	project *models.Project,
	repository repositories.ProjectRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	jobsRepository repositories.JobRepository,
	constructionJobMaterialRepository repositories.ConstructionJobMaterialRepository) dtos.Response {

	setProjectJobsResult := setProjectJobs(
		project, nil, constructionJobPropertyRepository, jobsRepository, constructionJobMaterialRepository)

	if setProjectJobsResult.Success {
		operationResult := repository.Save(setProjectJobsResult.Data.(*models.Project))

		if operationResult.Error != nil {
			return dtos.Response{Success: false, Message: operationResult.Error.Error()}
		}

		var data = operationResult.Result.(*models.Project)

		return dtos.Response{Success: true, Data: data}
	}

	return setProjectJobsResult
}

func UpdateProjectById(
	id string,
	project *models.Project,
	repository repositories.ProjectRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	projectJobRepository repositories.ProjectJobRepository,
	projectPropertyRepository repositories.ProjectPropertyRepository,
	jobsRepository repositories.JobRepository,
	constructionJobMaterialRepository repositories.ConstructionJobMaterialRepository,
	projectMaterialRepository repositories.ProjectMaterialRepository) dtos.Response {

	log.Println("Start update project with id: ", id)

	data, parseErr := base64.StdEncoding.DecodeString(project.ProjectCoverBase64)
	if parseErr != nil {
		log.Println("Project id:", id, "Error decoding image string:", parseErr)
	}
	project.ProjectCover = data

	projectPropertyDeleteResult := projectPropertyRepository.DeleteProjectPropertiesByProjectId(id)
	if projectPropertyDeleteResult.Error != nil {
		log.Println("Failed to remove properties for project with id: ", id)
		return dtos.Response{Success: false, Message: projectPropertyDeleteResult.Error.Error()}
	}
	log.Println("Succsessfully removed project properties for project with id: ", id)

	projectJobDeleteResult := projectJobRepository.DeleteProjectJobsByProjectId(id)
	if projectJobDeleteResult.Error != nil {
		log.Println("Failed to remove jobs for project with id: ", id)
		return dtos.Response{Success: false, Message: projectJobDeleteResult.Error.Error()}
	}

	log.Println("Succsessfully removed project jobs for project with id: ", id)

	projectMaterialDeleteResult := projectMaterialRepository.DeleteProjectMaterialsByProjectId(id)
	if projectMaterialDeleteResult.Error != nil {
		log.Println("Failed to remove jobs for project with id: ", id)
		return dtos.Response{Success: false, Message: projectMaterialDeleteResult.Error.Error()}
	}

	log.Println("Succsessfully removed project materials for project with id: ", id)

	existingProjectResponse := GetProjectById(id, repository)

	if !existingProjectResponse.Success {
		return existingProjectResponse
	}
	log.Println("Succsessfully retrieved project with id: ", id)
	existingProject := existingProjectResponse.Data.(*models.Project)

	existingProject.Name = project.Name
	existingProject.Filename = project.Filename
	existingProject.BucketName = project.BucketName
	existingProject.LivingArea = project.LivingArea
	existingProject.Margin = project.Margin
	existingProject.ProjectCover = project.ProjectCover
	existingProject.Workers = project.Workers
	existingProject.ConstructionCompanyName = project.ConstructionCompanyName
	existingProject.ConstructionWorkersNumber = project.ConstructionWorkersNumber
	existingProject.WallMaterial = project.WallMaterial
	existingProject.FoundationMaterial = project.FoundationMaterial
	existingProject.FinishMaterial = project.FinishMaterial
	existingProject.RoofingMaterial = project.RoofingMaterial
	existingProject.ProjectProperties = project.ProjectProperties

	log.Println("Succsessfully updated properties for project with id: ", id)

	setProjectJobsResult := setProjectJobs(
		project, existingProject, constructionJobPropertyRepository, jobsRepository, constructionJobMaterialRepository)

	if setProjectJobsResult.Success {
		operationResult := repository.Save(setProjectJobsResult.Data.(*models.Project))

		if operationResult.Error != nil {
			return dtos.Response{Success: false, Message: operationResult.Error.Error()}
		}

		var data = operationResult.Result.(*models.Project)
		data.ProjectCoverBase64 = base64.StdEncoding.EncodeToString(data.ProjectCover)

		return dtos.Response{Success: true, Data: data}
	}

	return setProjectJobsResult
}

func UpdateProjectProperties(
	id string,
	project *models.Project,
	repository repositories.ProjectRepository,
	jobsRepository repositories.JobRepository,
	projectJobRepository repositories.ProjectJobRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	constructionJobMaterialRepository repositories.ConstructionJobMaterialRepository,
	projectMaterialRepository repositories.ProjectMaterialRepository) dtos.Response {

	projectJobDeleteResult := projectJobRepository.DeleteProjectJobsByProjectId(id)
	if projectJobDeleteResult.Error != nil {
		log.Println("Failed to remove jobs for project with id: ", id)
		return dtos.Response{Success: false, Message: projectJobDeleteResult.Error.Error()}
	}

	log.Println("Succsessfully removed project jobs for project with id: ", id)

	projectMaterialDeleteResult := projectMaterialRepository.DeleteProjectMaterialsByProjectId(id)
	if projectMaterialDeleteResult.Error != nil {
		log.Println("Failed to remove jobs for project with id: ", id)
		return dtos.Response{Success: false, Message: projectMaterialDeleteResult.Error.Error()}
	}

	log.Println("Succsessfully removed project materials for project with id: ", id)

	log.Println("Start update project with id: ", id)

	existingProjectResponse := GetProjectById(id, repository)

	if !existingProjectResponse.Success {
		return existingProjectResponse
	}
	log.Println("Succsessfully retrieved project with id: ", id)
	existingProject := existingProjectResponse.Data.(*models.Project)

	existingProject.ConstructionWorkersNumber = project.ConstructionWorkersNumber
	existingProject.WallMaterial = project.WallMaterial
	existingProject.FoundationMaterial = project.FoundationMaterial
	existingProject.FinishMaterial = project.FinishMaterial
	existingProject.RoofingMaterial = project.RoofingMaterial

	log.Println("Succsessfully updated properties for project with id: ", id)

	setProjectJobsResult := setProjectJobs(
		project, existingProject, constructionJobPropertyRepository, jobsRepository, constructionJobMaterialRepository)

	if setProjectJobsResult.Success {
		operationResult := repository.Save(setProjectJobsResult.Data.(*models.Project))

		if operationResult.Error != nil {
			return dtos.Response{Success: false, Message: operationResult.Error.Error()}
		}

		var data = operationResult.Result.(*models.Project)

		return dtos.Response{Success: true, Data: data}
	}

	return setProjectJobsResult
}

func GetAllProjects(repository repositories.ProjectRepository) dtos.Response {
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*[]models.ProjectMin)
	var arr = *datas
	for index, data := range arr {
		arr[index].ProjectCoverBase64 = base64.StdEncoding.EncodeToString(data.ProjectCover)
	}

	return dtos.Response{Success: true, Data: datas}
}

func DeleteProjectById(
	id string, repository repositories.ProjectRepository,
	projectPropertyRepository repositories.ProjectPropertyRepository,
	projectJobRepository repositories.ProjectJobRepository,
	projectMaterialRepository repositories.ProjectMaterialRepository) dtos.Response {
	projectPropertyRepository.DeleteProjectPropertiesByProjectId(id)
	projectJobRepository.DeleteProjectJobsByProjectId(id)
	projectMaterialRepository.DeleteProjectMaterialsByProjectId(id)
	operationResult := repository.DeleteProjectById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dtos.Response{Success: true}
}

func GetProjectById(id string, repository repositories.ProjectRepository) dtos.Response {
	operationResult := repository.GetProjectById(id)

	if operationResult.Error != nil {
		return dtos.Response{Success: false, Message: operationResult.Error.Error()}
	}

	data := operationResult.Result.(*models.Project)

	data.ProjectCoverBase64 = base64.StdEncoding.EncodeToString(data.ProjectCover)

	return dtos.Response{Success: true, Data: operationResult.Result}
}

func UpdateProjectsJobs(
	projectRepository repositories.ProjectRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	projectJobRepository repositories.ProjectJobRepository,
	jobsRepository repositories.JobRepository,
	constructionJobMaterialRepository repositories.ConstructionJobMaterialRepository,
	projectMaterialsRepository repositories.ProjectMaterialRepository) dtos.Response {
	getAllProjectsResult := projectRepository.FindAllProjects()
	if getAllProjectsResult.Error != nil {
		log.Println("Failed to retrieve all project for update")
		return dtos.Response{Success: false, Message: getAllProjectsResult.Error.Error()}
	}

	// todo: replac3 with project's company name
	constructionPropertiesRepositoryResult := constructionJobPropertyRepository.
		FindPropertiesByCompanyName("Construction")

	if constructionPropertiesRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: constructionPropertiesRepositoryResult.Error.Error()}
	}

	projects := getAllProjectsResult.Result.(*[]models.Project)
	for _, project := range *projects {
		updateProjectJob(
			&project,
			&projectRepository,
			constructionJobPropertyRepository,
			projectJobRepository,
			jobsRepository,
			constructionJobMaterialRepository,
			projectMaterialsRepository)
	}

	return dtos.Response{Success: true}
}

func updateProjectJob(
	project *models.Project,
	repository *repositories.ProjectRepository,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	projectJobRepository repositories.ProjectJobRepository,
	jobsRepository repositories.JobRepository,
	constructionJobMaterialRepository repositories.ConstructionJobMaterialRepository,
	projectMaterialRepository repositories.ProjectMaterialRepository) dtos.Response {

	projectJobDeleteResult := projectJobRepository.DeleteProjectJobsByProjectId(strconv.Itoa(project.ProjectId))
	if projectJobDeleteResult.Error != nil {
		log.Println("Failed to remove jobs for project with id: ", project.ProjectId)
		return dtos.Response{Success: false, Message: projectJobDeleteResult.Error.Error()}
	}

	projectMaterialDeleteResult := projectMaterialRepository.DeleteProjectMaterialsByProjectId(strconv.Itoa(project.ProjectId))
	if projectMaterialDeleteResult.Error != nil {
		log.Println("Failed to remove materika for project with id: ", project.ProjectId)
		return dtos.Response{Success: false, Message: projectMaterialDeleteResult.Error.Error()}
	}

	log.Println("Succsessfully removed project materials for project with id: ", project.ProjectId)

	setProjectJobsResult := setProjectJobs(
		project, nil, constructionJobPropertyRepository, jobsRepository, constructionJobMaterialRepository)

	if setProjectJobsResult.Success {
		project.ProjectProperties = []models.ProjectProperty{}
		operationResult := repository.Save(setProjectJobsResult.Data.(*models.Project))

		if operationResult.Error != nil {
			return dtos.Response{Success: false, Message: operationResult.Error.Error()}
		}

		var data = operationResult.Result.(*models.Project)

		return dtos.Response{Success: true, Data: data}
	}

	return setProjectJobsResult
}

func setProjectJobs(
	project *models.Project,
	existingProject *models.Project,
	constructionJobPropertyRepository repositories.ConstructionJobPropertyRepository,
	jobsRepository repositories.JobRepository,
	constructionJobMaterialRepository repositories.ConstructionJobMaterialRepository) dtos.Response {
	constructionPropertiesRepositoryResult := constructionJobPropertyRepository.
		FindPropertiesByCompanyName(project.ConstructionCompanyName)

	if constructionPropertiesRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: constructionPropertiesRepositoryResult.Error.Error()}
	}

	projectJobsRepositoryResult := jobsRepository.FindProjectJobs(
		project.WallMaterial,
		project.FoundationMaterial,
		project.RoofingMaterial,
		project.FinishMaterial)
	if projectJobsRepositoryResult.Error != nil {
		return dtos.Response{Success: false, Message: projectJobsRepositoryResult.Error.Error()}
	}

	project.ProjectJobs = []models.ProjectJob{}
	project.ProjectMaterials = []models.ProjectJobMaterial{}
	if existingProject != nil {
		existingProject.ProjectJobs = []models.ProjectJob{}
		existingProject.ProjectMaterials = []models.ProjectJobMaterial{}
	}

	jobIds := []int{}
	for _, job := range *projectJobsRepositoryResult.Result.(*[]models.Job) {
		project.ProjectJobs = append(project.ProjectJobs, converters.ConvertJobToProjectJob(job))
		if existingProject != nil {
			existingProject.ProjectJobs = append(existingProject.ProjectJobs, converters.ConvertJobToProjectJob(job))
		}
		jobIds = append(jobIds, job.JobId)
	}

	projectMaterialsRepositoryResult := constructionJobMaterialRepository.FindProjectJobsMaterials(jobIds)

	projectJobsCalculated := helpers.CalculateCostDurationForProjectJobs(
		*project,
		constructionPropertiesRepositoryResult.Result.(*[]models.ConstructionJobProperty))

	calcJobMap := map[string]helpers.JobCalculations{}
	for _, prop := range projectJobsCalculated {
		calcJobMap[prop.JobCode] = prop
	}

	var projectToUpdate *models.Project
	if existingProject == nil {
		projectToUpdate = project
	} else {
		projectToUpdate = existingProject
	}

	propertiesMap := map[string]float32{}
	for _, p := range projectToUpdate.ProjectProperties {
		propertiesMap[p.PropertyCode] = p.PropertyValue
	}

	var materialsCost float32
	for _, material := range *projectMaterialsRepositoryResult.Result.(*[]models.ConstructionJobMaterial) {
		projectMaterial := converters.ConvertMaterialToProjectMaterial(material)
		projectMaterial.MaterialCost = propertiesMap[*material.Job.PropertyID] * material.MaterialCost
		materialsCost += projectMaterial.MaterialCost
		if existingProject != nil {
			projectMaterial.ProjectRefer = projectToUpdate.ProjectId
		}
		projectToUpdate.ProjectMaterials = append(projectToUpdate.ProjectMaterials, projectMaterial)
	}

	for index, job := range projectToUpdate.ProjectJobs {
		calJob := calcJobMap[job.Job.JobCode]
		(&projectToUpdate.ProjectJobs[index]).ConstructionCost = calJob.ConstructionCost
		(&projectToUpdate.ProjectJobs[index]).ConstructionDurationInDays = calJob.ConstructionDurationInDays
		(&projectToUpdate.ProjectJobs[index]).ConstructionDurationInHours = calJob.ConstructionDurationInHours
		(&projectToUpdate.ProjectJobs[index]).ConstructionWorkers = calJob.ConstructionWorkers
		(&projectToUpdate.ProjectJobs[index]).ConstructionDuration = calJob.ConstructionDuration
	}

	var projectCost int
	var projectDuration int
	for _, j := range projectJobsCalculated {
		projectCost += int(j.ConstructionCost)
		projectDuration += j.ConstructionDurationInDays
	}

	projectToUpdate.ConstructionCost = projectCost + int(materialsCost)
	projectToUpdate.ConstructionJobCost = projectCost
	projectToUpdate.ConstructionMaterialCost = int(materialsCost)
	projectToUpdate.ConstructionDuration = helpers.CalculateProjectDuration(projectToUpdate.ProjectJobs)
	log.Println(
		"Recalculated duration and cost for project with project_id ", project.ProjectId,
		", project cost: ", projectCost, "project duration: ", projectDuration)

	return dtos.Response{Success: true, Data: projectToUpdate}
}
