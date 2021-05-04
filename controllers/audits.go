package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
	"strconv"
)

func AllIncidentsLog(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	search := c.DefaultQuery("search", "")
	end := c.DefaultQuery("end", "")
	start := c.DefaultQuery("start", "")
	user := c.DefaultQuery("user", "")

	var incidents []m.AuditLogIncident
	var count int64

	if search != "" {
		db = db.Where("LIKE %?", search)
	}
	if user != "" {
		db = db.Where("user_id=?", user)
	}
	if start != "" && end != "" {
		db = db.Where("date BETWEEN CAST(? AS DATETIME) AND CAST(? AS DATETIME)", start, end)
	}

	db.Debug().Scopes(m.Pagination(page, limit)).Order("id desc").Preload("User").Preload("Location").Find(&incidents)
	db.Model(m.AuditLogIncident{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
	}
	paginator.Records = incidents
	c.JSON(200, paginator)
}
