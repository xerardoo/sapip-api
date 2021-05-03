package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"strings"
	"time"
)

var DB *gorm.DB

func InitDB() (db *gorm.DB) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8&parseTime=true", username, password, host, port)
	dsndb := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", username, password, host, port, database)

	db, err := gorm.Open(mysql.Open(dsndb), &gorm.Config{})
	if err != nil {
		if strings.Contains(err.Error(), "Unknown database") {
			db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				panic("DB ReOpen 1: " + err.Error())
			}
			err = CreateDB(db)
			if err != nil {
				panic("DB Create: " + err.Error())
			}
			db, err = gorm.Open(mysql.Open(dsndb), &gorm.Config{})
			if err != nil {
				panic("DB ReOpen 2: " + err.Error())
			}
		} else {
			panic("DB Connection:" + err.Error())
		}
	}

	// err = db.SetupJoinTable(&MenuItem{}, "Taxes", &MenuItemTax{})
	// if err != nil {
	// 	panic("DB SetupJoinTable: " + err.Error())
	// }
	err = db.AutoMigrate(
		&User{},
		&Patrol{},
		&Persona{},
		&Vehicle{},
		&Location{},
		&Incident{},
	)
	if err != nil {
		fmt.Println("DB Migration: ", err.Error())
		panic("DB Migration: " + err.Error())
	}

	InitIncidents(db)
	InitPersona(db)
	InitUser(db)
	return
}

func DropDB(db *gorm.DB) (err error) {
	err = db.Exec("DROP DATABASE " + os.Getenv("DB_DATABASE")).Error
	if err != nil {
		return
	}
	return
}

func CreateDB(db *gorm.DB) (err error) {
	err = db.Exec("CREATE DATABASE " + os.Getenv("DB_DATABASE")).Error
	if err != nil {
		return
	}
	return
}

func UseDB(db *gorm.DB, database string) (err error) {
	err = db.Exec("USE DATABASE " + database).Error
	if err != nil {
		return
	}
	return
}

func DateToMx(date string) (dt string, err error) {
	if date == "" {
		return
	}
	d := strings.Split(date, "T")
	tm, err := time.Parse("2006-01-02", d[0])
	if err != nil {
		return
	}
	dt = tm.Format("02-01-2006")
	return
}

func DateMxToSql(date string) (dt string, err error) {
	if date == "" {
		return
	}
	tm, err := time.Parse("02-01-2006", date)
	if err != nil {
		return
	}
	dt = tm.Format("2006-01-02")
	return
}
