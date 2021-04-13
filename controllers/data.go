package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
	"googlemaps.github.io/maps"
	"os"
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