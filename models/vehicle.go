package models

type Vehicle struct {
	Model
	// Marca, linea, modelo, color, placas de circulacion, serie,
	VIN          string `gorm:"size:250;not null;" sql:"index" json:"vin"`
	Type         string `gorm:"size:250;not null;" sql:"index" json:"type"` // responsable, afectado
	Brand        string `gorm:"size:250;not null;" sql:"index" json:"brand"`
	SubBrand     string `gorm:"size:250;not null;" sql:"index" json:"subbrand"`
	Modelo       string `gorm:"size:250;not null;" sql:"index" json:"modelo"`
	Color        string `gorm:"size:250;not null;" sql:"index" json:"color"`
	Plate        string `gorm:"size:50; not null;" sql:"index" json:"plate"`
	PlateState   string `gorm:"size:50; not null;" sql:"index" json:"plate_state"`
	Origin       string `gorm:"size:25; not null;" sql:"index" json:"origin"`
	UseType      string `gorm:"size:25; not null;" sql:"index" json:"use_type"`
	Observations string `gorm:"size:300; not null;" json:"observations"`
	Photo        string `gorm:"size:250;"  json:"photo"`
	// Fotos
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
