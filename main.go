package main

import (
	"fmt"
	"github.com/joho/godotenv"
	. "github.com/xerardoo/sapip/models"
	"github.com/xerardoo/sapip/routes"
	"log"
	"os"
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

	r := routes.Init()
	err = r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
}
