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
	"bitbucket.org/houmeteam/houme-go/models"
)

var identityKey = "id"

func SetupRoutes(
	corsConfigs *CorsConfigs,
	adminController *controllers.AdminController,
	projectsController *controllers.ProjectsController,
	constructionPropertiesController *controllers.ConstructionPropertiesController,
	commonController *controllers.CommonController,
	foremanController *controllers.ForemanController,
	companyController *controllers.CompanyController,
	companyJobController *controllers.CompanyJobController) *gin.Engine {
	route := gin.Default()

	route.Use(gin.Logger())

	domains := []string{corsConfigs.Domain}
	if corsConfigs.SubDomain != "" {
		domains = append(domains, corsConfigs.SubDomain)
	}
	route.Use(cors.New(cors.Config{
		AllowOrigins: domains,
		AllowMethods: []string{"POST", "OPTIONS", "GET", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Content-Length", "Content-Type", "Accept-Encoding",
			"X-CSRF-Token", "Authorization", "accept", "origin",
			"Cache-Control", "X-Requested-With", "Access-Control-Allow-Origin",
			"Cookie"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "test zone",
		Key:            []byte("secret key"),
		Timeout:        time.Hour,
		CookieMaxAge:   time.Hour,
		MaxRefresh:     time.Hour,
		IdentityKey:    identityKey,
		SendCookie:     true,
		CookieHTTPOnly: corsConfigs.IsProd,
		SecureCookie:   true,
		CookieSameSite: http.SameSiteNoneMode,
		CookieDomain:   corsConfigs.Domain,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			v, ok := data.(*models.LoginInfo)
			if ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"Role":      v.Role,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.LoginInfo{
				Email: claims[identityKey].(string),
				Role:  claims["Role"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.LoginInfo
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			response := adminController.LoginAdminHandler(loginVals)
			if response.Success {
				return &models.LoginInfo{
					Email: loginVals.Email,
					Role:  "admin",
				}, nil
			} else {
				response := foremanController.LoginForemanHandler(&loginVals)
				if response.Success {
					return &models.LoginInfo{
						Email: loginVals.Email,
						Role:  "foreman",
					}, nil
				}
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.LoginInfo); ok && v.Role == "admin" || v.Role == "foreman" {
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
		TokenLookup: "cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

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
	//route.POST("/registerAdmin", adminController.RegisterAdminHandler)

	route.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := route.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.POST("/logout", authMiddleware.LogoutHandler)

	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/getUserInfo", adminController.GetUserInfoHandler)
		auth.POST("/create", projectsController.CreateProjectHandler)
		auth.PUT("/updateProject/:id", projectsController.UpdateProjectByIdHandler)
		auth.DELETE("/deleteProject/:id", projectsController.DeleteProjectByIdHandler)
		auth.GET("/getProjectJobs/:projectId", projectsController.GetProjectJobsByProjectIdHandler)
		auth.PUT("/updateProjects", projectsController.UpdateProjectsHandler)

		auth.GET("/getJobProperties", constructionPropertiesController.GetJobPropertiesHandler)
		auth.PUT("/updateJobProperty/:id", constructionPropertiesController.UpdateJobPropertyByIdHandler)

		auth.POST("/createMaterial", constructionPropertiesController.CreateMaterialHandler)
		auth.DELETE("/deleteJobMaterial/:id", constructionPropertiesController.DeleteJobMaterialByIdHandler)
		auth.PUT("/updateJobMaterial/:id", constructionPropertiesController.UpdateJobMaterialByIdHandler)
		auth.GET("/getJobMaterials", constructionPropertiesController.GetMaterialsHandler)

		auth.POST("/upload", commonController.ForgeUploadHandler)
		auth.POST("/translate", commonController.ForgeTranslateHandler)
		auth.POST("/createBucket", func(c *gin.Context) {
			forge.CreateBucket("houmly")
		})
		auth.GET("/forgeGet", commonController.ForgeGetHandler)
		auth.GET("/translationStatus", commonController.ForgeTranslationStatusHandler)

		auth.POST("/company/create", companyController.AddCompanyHandler)
		auth.GET("/company/get", companyController.GetCompaniesHandler)
		auth.DELETE("/company/delete/:id", companyController.DeleteCompanyHandler)
	}

	foreman := route.Group("/foreman")
	foreman.Use(authMiddleware.MiddlewareFunc())
	{
		auth.POST("/getJobs/:companyId", companyJobController.GetCompanyJobsHandler)
		auth.DELETE("/deleteJob/:id", companyJobController.DeleteCompanyJobHandler)
		auth.PUT("/updateJob/:id", companyJobController.UpdateCompanyJobHandler)
		auth.GET("/createJob/:id", companyJobController.CreateCompanyJobHandler)
	}

	route.GET("/getProjects", projectsController.GetProjectsHandler)

	route.GET("/getProperties", commonController.GetProperties)

	route.GET("/getJobs", commonController.GetJobsHandler)

	route.PUT("/updateProjectProperties/:id", projectsController.UpdateProjectPropertiesByIdHandler)

	route.GET("/getProject/:id", projectsController.GetProjectByIdHandler)

	return route
}
