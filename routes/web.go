package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xerardoo/sapip/controllers"
	"time"
)

func Init() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())

	auth := r.Group("/v1")
	{
		auth.POST("/signin", controllers.Signin)
		auth.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"time": time.Now(),
				"utc":  time.UTC.String(),
			})
		})
	}

	data := r.Group("/v1/data")
	{
		data.POST("/geocodingr", controllers.GetGeocodingReverse)
		data.GET("/incident-types", controllers.AllIncidentTypes)
		data.GET("/persona-types", controllers.AllPersonaTypes)
	}
	data.Use(VerifyJWT())

	incident := r.Group("/v1/incident")
	{
		incident.GET("", controllers.AllIncidents)
		incident.GET("/:id", controllers.FindIncident)
		incident.POST("", controllers.AddIncident)
		// incident.PUT("/:id", controllers.UpdIncident)
		// incident.DELETE("/:id", controllers.DelIncident)
	}
	incident.Use(VerifyJWT())

	incidentAdmin := r.Group("/v1/admin/incident")
	{
		incidentAdmin.GET("", controllers.AllIncidents)
		incidentAdmin.GET("/:id", controllers.FindIncident)
	}

	userAdmin := r.Group("/v1/admin/user")
	{
		userAdmin.GET("", controllers.AllUsers)
		userAdmin.POST("", controllers.AddUser)
		userAdmin.GET("/:id", controllers.FindUser)
		userAdmin.PUT("/:id", controllers.UpdUser)
		userAdmin.DELETE("/:id", controllers.DelUser)
	}

	dataAdmin := r.Group("/v1/admin/data")
	{
		dataAdmin.POST("/geocodingr", controllers.GetGeocodingReverse)
		dataAdmin.GET("/incident-types", controllers.AllIncidentTypes)
		dataAdmin.GET("/persona-types", controllers.AllPersonaTypes)
	}

	return r
}
