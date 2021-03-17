package tester

import (
	"github.com/joho/godotenv"
	"github.com/xerardoo/sapip/models"
	"os"
)

var envpath = "../.env"

func init() {
	if _, err := os.Stat(envpath); os.IsNotExist(err) {
		envpath = "../../.env"
	}
	err := godotenv.Load(envpath)
	if err != nil {
		panic(err)
	}
	models.DB = models.InitDB()
}