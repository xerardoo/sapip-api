package controllers

import (
	"github.com/gin-gonic/gin"
	m "github.com/xerardoo/sapip/models"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signin(c *gin.Context) {
	var creds Credentials
	user := new(m.User)

	err := c.BindJSON(&creds)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}
	if creds.Password == "" || len(creds.Password) < 4 {
		c.JSON(400, gin.H{"msg": "credenciales invalidas"})
		return
	}

	user, err = user.VerifyCredentials(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	token, exp, err := user.GenerateTokenJWT()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: exp,
	})
	c.JSON(200, gin.H{"user": user, "token": token})
}

func SigninAdmin(c *gin.Context) {
	var creds Credentials
	user := new(m.User)

	err := c.BindJSON(&creds)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}
	if creds.Password == "" || len(creds.Password) < 4 {
		c.JSON(400, gin.H{"msg": "credenciales invalidas"})
		return
	}

	user, err = user.VerifyCredentials(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": err.Error()})
		return
	}

	token, exp, err := user.GenerateTokenJWT()
	if err != nil {
		c.JSON(500, gin.H{"msg": err.Error()})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: exp,
	})
	c.JSON(200, gin.H{"user": user, "token": token})
}
