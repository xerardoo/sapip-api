package models

import (
	"time"
)

type User struct {
	Model
	FirstName   string `gorm:"size:250;not null;" sql:"index" json:"first_name"`
	LastName    string `gorm:"size:250;not null;" sql:"index" json:"last_name"`
	Password    string `gorm:"size:250;not null;" sql:"index" json:"-"`
	Rank        string `gorm:"size:250;not null;" sql:"index" json:"rank"`
	PhoneNumber string `gorm:"size:250;not null;" sql:"index" json:"phone_number"`
	Email       string `gorm:"size:250;" json:"email"`
	Birthday    time.Time
	ActivatedAt time.Time
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
	err = DB.Save(&u).Error
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
