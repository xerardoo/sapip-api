package models

type Vehicle struct {
	Model
	VIN        string `gorm:"size:250;not null;" sql:"index" json:"vin"`
	Type       string `gorm:"size:250;not null;" sql:"index" json:"type"`
	Brand      string `gorm:"size:250;not null;" sql:"index" json:"brand"`
	Modelo     string `gorm:"size:250;not null;" sql:"index" json:"modelo"`
	Year       string `gorm:"size:250;not null;" sql:"index" json:"year"`
	Color      string `gorm:"size:250;not null;" sql:"index" json:"color"`
	Plate      string `gorm:"size:50; not null;" sql:"index" json:"plate"`
	PlateState string `gorm:"size:50; not null;" sql:"index" json:"plate_state"`
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
