package main

import (
	"github.com/joho/godotenv"
	"os"
	. "github.com/xerardoo/sapip/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dir, _ := os.Getwd()
	err = os.Setenv("CURR_DIR", dir)
	if err != nil {
		panic(err)
	}

	err = os.Setenv("GIN_MODE", os.Getenv("APP_ENV"))
	if err != nil {
		panic(err)
	}

	DB = InitDB()
	// defer DB.Close()
}