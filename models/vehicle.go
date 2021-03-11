package models

type Vehicle struct {
	Model
	VIN    string `gorm:"size:250;not null;" sql:"index"`
	Brand  string `gorm:"size:250;not null;" sql:"index"`
	Modelo string `gorm:"size:250;not null;" sql:"index"`
	Year   string `gorm:"size:250;not null;" sql:"index"`
	Color  string `gorm:"size:250;not null;" sql:"index"`
}
