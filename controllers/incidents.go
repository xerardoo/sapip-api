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
	search := c.DefaultQuery("search", "")
	end := c.DefaultQuery("end", "")
	start := c.DefaultQuery("start", "")
	user := c.DefaultQuery("user", "")
	typee := c.DefaultQuery("type", "")
	cp := c.DefaultQuery("cp", "")

	// sortBy := c.DefaultQuery("sortBy", "id")
	// order := c.DefaultQuery("order", "desc")

	var incidents []m.Incident
	var count int64

	if search != "" {
		db = db.Where("MATCH(description) AGAINST (? IN NATURAL LANGUAGE MODE)", search)
	}
	if user != "" {
		db = db.Where("user_id=?", user)
	}
	if typee != "" {
		db = db.Where("type_id=?", typee)
	}
	if cp != "" {
		db = db.Where("zip_code=?", cp)
	}
	if start != "" && end != "" {
		db = db.Where("date BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)", start, end)
	}

	db.Debug().Scopes(m.Pagination(page, limit)).Order("id desc").Preload("Type").Preload("User").Find(&incidents)
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
		personasCount, err := incident.GetPersonasCount()
		if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		vehiclesCount, err := incident.GetVehiclesCount()
		if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
			c.JSON(500, gin.H{"msg": err.Error()})
			return
		}
		// user, err := incident.GetUser()
		// if err != nil && gorm.ErrRecordNotFound.Error() != err.Error() {
		// 	c.JSON(500, gin.H{"msg": err.Error()})
		// 	return
		// }
		// date, err := m.DateToMx(incident.Date)
		// if err != nil {
		// 	c.JSON(500, gin.H{"msg": err.Error()})
		// 	return
		// }

		// incidents[i].Date = date
		incidents[i].Location = location
		incidents[i].PersonasCount = personasCount
		incidents[i].VehiclesCount = vehiclesCount
		// incidents[i].User = user
	}
	// if len(incidents)==0{
	// 	incidents = []
	// }
	paginator.Records = incidents
	c.JSON(200, paginator)
}

func FindIncident(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userData, _ := c.Get("USER")
	user := userData.(*m.User)
	ua := c.Request.Header.Get("User-Agent")

	var incident m.Incident
	err := incident.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	for i, p := range incident.Personas {
		incident.Personas[i].PhotoFront = m.GetFilePathS3(p.PhotoFront)
	}
	for i, v := range incident.Vehicles {
		incident.Vehicles[i].Photo = m.GetFilePathS3(v.Photo)
	}

	audit := m.AuditLogIncident{Action: m.AUDIT_INCIDENT_VISIT, UserAgent: ua, IncidentID: incident.ID, UserID: user.ID}
	_, err = audit.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, incident)
}

func AddIncident(c *gin.Context) {
	var incident m.Incident
	userData, _ := c.Get("USER")
	user := userData.(*m.User)

	err := c.BindJSON(&incident)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	incident.Date, err = m.DateMxToSql(incident.Date)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	incident.UserID = user.ID
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
