package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
	"googlemaps.github.io/maps"
	"os"
	"strconv"
)

func GetGeocodingReverse(c *gin.Context) {
	var location m.Location
	err := c.BindJSON(&location)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	client, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GOOGLE_MAPS_API_KEY")))
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	r := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{Lat: location.Latitude, Lng: location.Longitude},
		Region: "mx",
	}
	result, err := client.ReverseGeocode(context.Background(), r)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
	}
	c.JSON(200, result)
}

func AllIncidentTypes(c *gin.Context) {
	var incidents []m.IncidentType
	m.DB.Order("name desc").Find(&incidents)
	c.JSON(200, incidents)
}

func AllPersonaTypes(c *gin.Context) {
	var personas []m.PersonaType
	m.DB.Order("name desc").Find(&personas)
	c.JSON(200, personas)
}

func AddAuditEvent(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userData, _ := c.Get("USER")
	user := userData.(*m.User)
	ua := c.Request.Header.Get("User-Agent")

	audit := m.AuditLogIncident{Action: m.AUDIT_INCIDENT_COPY, UserAgent: ua, IncidentID: id, UserID: user.ID}
	_, err := audit.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.Status(200)
}
