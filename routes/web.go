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

	authAdmin := r.Group("/v1/admin")
	{
		authAdmin.POST("/signin", controllers.SigninAdmin)
		authAdmin.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"time": time.Now(),
				"utc":  time.UTC.String(),
			})
		})
	}

	data := r.Group("/v1/data").Use(VerifyJWT())
	{
		data.POST("/geocodingr", controllers.GetGeocodingReverse)
		data.GET("/incident-types", controllers.AllIncidentTypes)
		data.GET("/persona-types", controllers.AllPersonaTypes)
	}

	meta := r.Group("/v1/meta").Use(VerifyJWT())
	{
		meta.POST("/session/:id", controllers.AddAuditEvent)
	}

	incident := r.Group("/v1/incident").Use(VerifyJWT())
	{
		incident.GET("", controllers.AllIncidents)
		incident.GET("/:id", controllers.FindIncident)
		incident.POST("", controllers.AddIncident)
		// incident.PUT("/:id", controllers.UpdIncident)
		// incident.DELETE("/:id", controllers.DelIncident)
	}

	incidentAdmin := r.Group("/v1/admin/incident").Use(VerifyJWT())
	{
		incidentAdmin.GET("", controllers.AllIncidents)
		incidentAdmin.GET("/:id", controllers.FindIncident)
	}

	userAdmin := r.Group("/v1/admin/user").Use(VerifyJWT())
	{
		userAdmin.GET("", controllers.AllUsers)
		userAdmin.POST("", controllers.AddUser)
		userAdmin.GET("/:id", controllers.FindUser)
		userAdmin.PUT("/:id", controllers.UpdUser)
		userAdmin.DELETE("/:id", controllers.DelUser)
	}

	dataAdmin := r.Group("/v1/admin/data").Use(VerifyJWT())
	{
		dataAdmin.POST("/geocodingr", controllers.GetGeocodingReverse)
		dataAdmin.GET("/incident-types", controllers.AllIncidentTypes)
		dataAdmin.GET("/persona-types", controllers.AllPersonaTypes)
	}

	auditLog := r.Group("/v1/admin/audit-log").Use(VerifyJWT())
	{
		auditLog.GET("/incident", controllers.AllIncidentsLog)
	}

	return r
}
