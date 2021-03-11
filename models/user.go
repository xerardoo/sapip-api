package models

import (
	"database/sql"
	"time"
)

type User struct {
	Model
	FirstName   string `gorm:"size:250;not null;" sql:"index"`
	LastName    string `gorm:"size:250;not null;" sql:"index"`
	Password    string `gorm:"size:250;not null;" sql:"index"`
	Rank        string `gorm:"size:250;not null;" sql:"index"`
	PhoneNumber string `gorm:"size:250;not null;" sql:"index"`
	Email       string `gorm:"size:250;"`
	Birthday    time.Time
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
