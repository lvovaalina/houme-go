package configs

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"

	"bitbucket.org/houmeteam/houme-go/controllers"
	"bitbucket.org/houmeteam/houme-go/forge"
	"bitbucket.org/houmeteam/houme-go/helpers"
	"bitbucket.org/houmeteam/houme-go/models"
	"bitbucket.org/houmeteam/houme-go/repositories"
	"bitbucket.org/houmeteam/houme-go/services"
)

var identityKey = "id"

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*models.Admin).Email,
		"text":     "Hello World.",
	})
}

func SetupRoutes(
	corsConfigs *CorsConfigs,
	adminController *controllers.AdminController,
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
		AllowOrigins: []string{corsConfigs.Domain},
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Content-Length", "Content-Type", "Accept-Encoding",
			"X-CSRF-Token", "Authorization", "accept", "origin",
			"Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin",
			"Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	var identityKey = "id"
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "test zone",
		Key:            []byte("secret key"),
		Timeout:        time.Hour,
		MaxRefresh:     time.Hour,
		IdentityKey:    identityKey,
		SendCookie:     true,
		CookieHTTPOnly: corsConfigs.IsProd,
		SecureCookie:   corsConfigs.IsProd,
		CookieDomain:   corsConfigs.Domain,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			v, ok := data.(*models.Admin)
			log.Println("IS OK", ok)
			if ok {
				log.Println(v.Email)
				return jwt.MapClaims{
					identityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.Admin{
				Email: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.Admin
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if (loginVals.Email == "admin" && loginVals.PasswordString == "admin") ||
				(loginVals.Email == "test" && loginVals.PasswordString == "test") {
				return &models.Admin{
					Email: loginVals.Email,
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			log.Println(data)
			if v, ok := data.(*models.Admin); ok && v.Email == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	route.POST("/login", authMiddleware.LoginHandler)

	route.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := route.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
	}

	route.POST("/registerAdmin", adminController.RegisterAdminHandler)

	route.PUT("/loginAdmin", adminController.LoginAdminHandler)

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
