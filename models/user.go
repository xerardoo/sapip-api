package models

import (
	"crypto/md5"
	"encoding/hex"
	"gorm.io/gorm"
)

type User struct {
	Model
	FirstName   string `gorm:"size:250;not null;" sql:"index" json:"first_name"`
	LastName    string `gorm:"size:250;not null;" sql:"index" json:"last_name"`
	BadgeNumber string `gorm:"size:250;not null;" sql:"index" json:"badge_number"`
	Password    string `gorm:"size:250;not null;" sql:"index" json:"-"`
	Password1   string `gorm:"-" json:"password_1"`
	Password2   string `gorm:"-" json:"password_2"`
	Rank        string `gorm:"size:250;not null;" sql:"index" json:"rank"`
	// PhoneNumber string    `gorm:"size:250;not null;" sql:"index" json:"phone_number"`
	// Email       string    `gorm:"size:250;" json:"email"`
	// LastLogin   time.Time `json:"last_login"`
	// BirthDay    time.Time `json:"birth_day"`
	// ActivatedAt time.Time
}

func InitUser(db *gorm.DB) {
	var password = "wk9LAPXtUA2a6UqETws"
	hashPassword := md5.Sum([]byte(password))
	db.FirstOrCreate(&User{Model: Model{ID: 1}, BadgeNumber: "10000", FirstName: "System Admin", Password: hex.EncodeToString(hashPassword[:])})
}

func (u *User) Add() (*User, error) {
	hashPassword := md5.Sum([]byte(u.Password))
	u.Password = hex.EncodeToString(hashPassword[:])
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
	if u.Password != "" {
		hashPassword := md5.Sum([]byte(u.Password))
		u.Password = hex.EncodeToString(hashPassword[:])
	} else {
		// TODO restore password
		u.Password = uu.Password
	}
	err = DB.Save(&u).Error
	if err != nil {
		return nil, err
	}
	return u, err
}

func (u *User) Remove() (err error) {
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

func (u *User) VerifyCredentials(user string, password string) (pu *User, err error) {
	hashPassword := md5.Sum([]byte(password))
	err = DB.First(&u,"badge_number=? AND password=?", user, hex.EncodeToString(hashPassword[:])).Error
	if err != nil {
		return
	}
	pu = u // fix, copy after reassign
	return
}
