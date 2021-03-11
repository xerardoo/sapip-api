package models

type Location struct {
	Model
	Latitude  string `gorm:"size:32"`
	Longitude string `gorm:"size:32"`
}
