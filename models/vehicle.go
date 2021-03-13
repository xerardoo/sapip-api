package models

type Vehicle struct {
	Model
	VIN        string `gorm:"size:250;not null;" sql:"index"`
	Type       string `gorm:"size:250;not null;" sql:"index"`
	Brand      string `gorm:"size:250;not null;" sql:"index"`
	Modelo     string `gorm:"size:250;not null;" sql:"index"`
	Year       string `gorm:"size:250;not null;" sql:"index"`
	Color      string `gorm:"size:250;not null;" sql:"index"`
	Plate      string `gorm:"size:50; not null;" sql:"index"`
	PlateState string `gorm:"size:50; not null;" sql:"index"`
}

func (v *Vehicle) Add() (*Vehicle, error) {
	err := DB.Create(&v).Error
	if err != nil {
		return nil, err
	}
	return v, err
}

func (v *Vehicle) Find(id int) (err error) {
	err = DB.First(&v, id).Error
	if err != nil {
		return
	}
	return
}

func (v *Vehicle) Update() (*Vehicle, error) {
	var vv Vehicle
	err := DB.First(&vv, v.ID).Error
	if err != nil {
		return nil, err
	}
	err = DB.Save(&v).Error
	if err != nil {
		return nil, err
	}
	return v, err
}

func (v *Vehicle) Remove() (err error) {
	err = DB.First(&v, v.ID).Error
	if err != nil {
		return
	}
	err = DB.Delete(&v).Error
	if err != nil {
		return
	}
	return
}
