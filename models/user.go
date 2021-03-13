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

func (u *User) Add() (*User, error) {
	err := DB.Create(&u).Error
	if err != nil {
		return nil, err
	}
	return u, err
}

func (u *User) Find(id int) (err error) {
	err = DB.First(&u, id).Error
	if err != nil {
		return
	}
	return
}

func (u *User) Update() (*User, error) {
	var uu User
	err := DB.First(&uu, u.ID).Error
	if err != nil {
		return nil, err
	}
	err = DB.Saue(&u).Error
	if err != nil {
		return nil, err
	}
	return u, err
}

func (u *User) Remoue() (err error) {
	err = DB.First(&u, u.ID).Error
	if err != nil {
		return
	}
	err = DB.Delete(&u).Error
	if err != nil {
		return
	}
	return
}
