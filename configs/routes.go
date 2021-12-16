package configs

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"bitbucket.org/houmeteam/houme-go/forge"
	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
)

func SetupRoutes(
	projectRepository *repositories.ProjectRepository,
	propertiesRepository *repositories.PropertyRepository,
	jobsRepository *repositories.JobRepository,
	constructionJobPropertiesRepository *repositories.ConstructionJobPropertyRepository,
	constructionJobMaterialsRepository *repositories.ConstructionJobMaterialRepository,
	projectJobRepository *repositories.ProjectJobRepository,
	projectPropertyRepository *repositories.ProjectPropertyRepository,
	projectMaterialRepository *repositories.ProjectMaterialRepository,
	adminRepository *repositories.AdminRepository) *gin.Engine {
	route := gin.Default()

	route.Use(gin.Logger())

	route.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://houmly-dev.herokuapp.com"},
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Content-Length", "Content-Type", "Accept-Encoding",
			"X-CSRF-Token", "Authorization", "accept", "origin",
			"Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	route.POST("/registerAdmin", func(context *gin.Context) {
		var admin models.Admin

		err := context.ShouldBindJSON(&admin)

		// validation errors
		if err != nil {
			log.Println("Cannot unmarshal admin data, error: ", err.Error())
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		// save project & get it's response
		response := services.AdminRegister(*adminRepository, admin)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/loginAdmin", func(context *gin.Context) {
		var admin models.Admin

		err := context.ShouldBindJSON(&admin)

		// validation errors
		if err != nil {
			log.Println("Cannot unmarshal admin data, error: ", err.Error())
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK
		response, token := services.AdminLogin(*adminRepository, admin)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.SetCookie("jwt", token, 60*60*2, "/", "https://houmly-dev.herokuapp.com", false, true)
		context.Writer.Header().Add("access-control-expose-headers", "Set-Cookie")
		context.JSON(code, response)
	})

	route.GET("/isLoggedIn", func(context *gin.Context) {
		authCookie, err := context.Cookie("jwt")
		if err != nil {
			log.Println("ERROR: ", err.Error())
			response := err.Error()

			context.JSON(http.StatusBadRequest, response)

			return
		}

		log.Println(authCookie)

		isAuthentificated, _ := services.IsAuthentificated(authCookie)
		if !isAuthentificated {
			context.JSON(http.StatusUnauthorized, "Not authorized")
			return
		}

		context.JSON(http.StatusOK, "logged in")

		return
	})
	route.POST("/create", func(context *gin.Context) {
		// initialization project model
		var project models.Project

		// validate json
		err := context.ShouldBindJSON(&project)

		// validation errors
		if err != nil {
			log.Println("Cannot unmarshal project, error: ", err.Error())
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save project & get it's response
		response := services.CreateProject(
			&project, *projectRepository, *constructionJobPropertiesRepository,
			*jobsRepository, *constructionJobMaterialsRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getProperties", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllProperties(*propertiesRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getJobs", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindAllJobs(*jobsRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getProjects", func(context *gin.Context) {
		code := http.StatusOK

		response := services.GetAllProjects(*projectRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/deleteProject/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteProjectById(
			id, *projectRepository, *projectPropertyRepository, *projectJobRepository, *projectMaterialRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getProjectJobs/:projectId", func(context *gin.Context) {
		projectId := context.Param("projectId")

		code := http.StatusOK

		response := services.FindJobsByProjectId(
			projectId,
			*projectJobRepository,
			*projectPropertyRepository,
			*constructionJobPropertiesRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getProject/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.GetProjectById(id, *projectRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/updateProject/:id", func(context *gin.Context) {
		id := context.Param("id")

		var project models.Project

		err := context.ShouldBindJSON(&project)

		// validation errors
		if err != nil {
			log.Println("ERROR: ", err.Error())
			response := err.Error()

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateProjectById(
			id, &project, *projectRepository,
			*constructionJobPropertiesRepository,
			*projectJobRepository, *projectPropertyRepository,
			*jobsRepository, *constructionJobMaterialsRepository,
			*projectMaterialRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/updateProjectProperties/:id", func(context *gin.Context) {
		id := context.Param("id")

		var project models.Project

		err := context.ShouldBindJSON(&project)

		// validation errors
		if err != nil {
			log.Println("ERROR: ", err.Error())
			response := err.Error()

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateProjectProperties(
			id, &project, *projectRepository,
			*jobsRepository, *projectJobRepository,
			*constructionJobPropertiesRepository,
			*constructionJobMaterialsRepository, *projectMaterialRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getJobProperties", func(context *gin.Context) {
		code := http.StatusOK

		authCookie, err := context.Cookie("jwt")
		if err != nil {
			log.Println("ERROR: ", err.Error())
			response := err.Error()

			context.JSON(http.StatusBadRequest, response)

			return
		}

		log.Println(authCookie)

		isAuthentificated, _ := services.IsAuthentificated(authCookie)
		if !isAuthentificated {
			context.JSON(http.StatusUnauthorized, "Not authorized")
			return
		}

		response := services.FindJobProperties(*constructionJobPropertiesRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/updateJobProperty/:id", func(context *gin.Context) {
		id := context.Param("id")

		var jobProperty models.ConstructionJobProperty

		err := context.ShouldBindJSON(&jobProperty)

		// validation errors
		if err != nil {
			log.Println(err.Error())
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateJobPropertyById(
			id, jobProperty, *constructionJobPropertiesRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/getJobMaterials", func(context *gin.Context) {
		code := http.StatusOK

		response := services.FindJobMaterials(*constructionJobMaterialsRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/updateJobMaterial/:id", func(context *gin.Context) {
		id := context.Param("id")

		var jobMaterial models.ConstructionJobMaterial

		err := context.ShouldBindJSON(&jobMaterial)

		// validation errors
		if err != nil {
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		code := http.StatusOK

		response := services.UpdateJobMaterialById(
			id, jobMaterial, *constructionJobMaterialsRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.DELETE("/deleteJobMaterial/:id", func(context *gin.Context) {
		id := context.Param("id")

		code := http.StatusOK

		response := services.DeleteJobMaterialById(id, *constructionJobMaterialsRepository)

		if !response.Success {
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.POST("/createMaterial", func(context *gin.Context) {
		var material models.ConstructionJobMaterial

		// validate json
		err := context.ShouldBindJSON(&material)

		// validation errors
		if err != nil {
			log.Println("Cannot unmarshal project, error: ", err.Error())
			// generate validation errors response
			response := helpers.GenerateValidationResponse(err)

			context.JSON(http.StatusBadRequest, response)

			return
		}

		// default http status code = 200
		code := http.StatusOK

		// save project & get it's response
		response := services.CreateMaterial(&material, *constructionJobMaterialsRepository)

		// save contact failed
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.PUT("/updateProjects", func(context *gin.Context) {
		code := http.StatusOK

		response := services.UpdateProjectsJobs(
			*projectRepository,
			*constructionJobPropertiesRepository,
			*projectJobRepository,
			*jobsRepository,
			*constructionJobMaterialsRepository,
			*projectMaterialRepository)
		if !response.Success {
			// change http status code to 400
			code = http.StatusBadRequest
		}

		context.JSON(code, response)
	})

	route.GET("/forgeGet", func(context *gin.Context) {
		projects := forge.GetBucketObjects("houme")
		context.JSON(http.StatusOK, projects)
	})

	return route
}
