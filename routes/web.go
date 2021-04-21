package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/xerardoo/sapip/controllers"
)

func Init() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(CORSMiddleware())

	// auth := r.Group("/v1")
	// {
	// 	auth.Any("/signin", controllers.Signin)
	// 	auth.Any("/logout", controllers.Logout)
	// 	auth.GET("/hello", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"time": time.Now(),
	// 			"utc":  time.UTC.String(),
	// 		})
	// 	})
	// }

	data := r.Group("/v1/data")
	{
		data.POST("/geocodingr", controllers.GetGeocodingReverse)
		data.GET("/incident-types", controllers.AllIncidentTypes)
		data.GET("/persona-types", controllers.AllPersonaTypes)
	}

	incident := r.Group("/v1/incident")
	{
		incident.GET("", controllers.AllIncidents)
		incident.GET("/:id", controllers.FindIncident)
		incident.POST("", controllers.AddIncident)
		// incident.PUT("/:id", controllers.UpdIncident)
		// incident.DELETE("/:id", controllers.DelIncident)
	}

	return r
}
