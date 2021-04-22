package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
	"gorm.io/gorm"
	"strconv"
)

func AllIncidents(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	// name := c.DefaultQuery("name", "")
	// sortBy := c.DefaultQuery("sortBy", "id")
	// order := c.DefaultQuery("order", "desc")

	var incidents []m.Incident
	var count int64
	// db = db.Where("name LIKE ?", name+"%")
	db.Scopes(m.Pagination(page, limit)).Order("id desc").Find(&incidents)
	db.Model(m.Incident{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
	}
	for i, incident := range incidents {
		location, err := incident.GetLocation()
		if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		// user, err := incident.GetUser()
		// if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
		// 	c.JSON(500, gin.H{"msg": err.Error()})
		// 	return
		// }
		incidents[i].Location = location
		// incidents[i].User = user
	}
	paginator.Records = incidents
	c.JSON(200, paginator)
}

func FindIncident(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var incident m.Incident
	err := incident.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	location, err := incident.GetLocation()
	if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	// user, err := incident.GetUser()
	// if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
	// 	c.JSON(500, gin.H{"msg": err.Error()})
	// 	return
	// }
	incident.Location = location
	// incident.User = user
	c.JSON(200, incident)
}

func AddIncident(c *gin.Context) {
	var incident m.Incident
	err := c.BindJSON(&incident)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	newincident, err := incident.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, newincident)
}

func UpdIncident(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var incident m.Incident
	err := c.BindJSON(&incident)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	incident.ID = id
	newincident, err := incident.Update()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, newincident)
}

func DelIncident(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var incident m.Incident
	incident.ID = id
	err := incident.Remove()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
