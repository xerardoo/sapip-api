package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
	"strconv"
)

func AllUsers(c *gin.Context) {
	db := m.DB
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	search := c.DefaultQuery("search", "")

	// sortBy := c.DefaultQuery("sortBy", "id")
	// order := c.DefaultQuery("order", "desc")

	var users []m.User
	var count int64

	if search != "" {
		db = db.Where("", search)
	}

	db.Debug().Scopes(m.Pagination(page, limit)).Order("id desc").Find(&users)
	db.Model(m.User{}).Count(&count)
	paginator := m.Paginator{
		Limit:       limit,
		Page:        page,
		TotalRecord: count,
	}
	paginator.Records = users
	c.JSON(200, paginator)
}

func FindUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user m.User
	err := user.Find(id)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, user)
}

func AddUser(c *gin.Context) {
	var user m.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	if user.Password1 != user.Password2 {
		c.JSON(400, gin.H{"msg": "Contraseñas no coinciden"})
		return
	}

	user.Password = user.Password1
	newUser, err := user.Add()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(201, newUser)
}

func UpdUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user m.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	if user.Password1 != user.Password2 {
		c.JSON(400, gin.H{"msg": "Contraseñas no coinciden"})
		return
	}

	user.ID = id
	user.Password = user.Password1
	newUser, err := user.Update()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, newUser)
}

func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user m.User
	user.ID = id
	err := user.Remove()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(200, gin.H{"msg": "ok"})
}
