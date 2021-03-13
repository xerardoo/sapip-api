package models

type Location struct {
	Model
	Latitude  string `gorm:"size:32"`
	Longitude string `gorm:"size:32"`
	TimeStamp string `gorm:"size:36"`
	Accuracy  string `gorm:"size:36"`
}

func (l *Location) Add() (*Location, error) {
	err := DB.Create(&l).Error
	if err != nil {
		return nil, err
	}
	return l, err
}

func (l *Location) Find(id int) (err error) {
	err = DB.First(&l, id).Error
	if err != nil {
		return
	}
	return
}

func (l *Location) Update() (*Location, error) {
	var ll Location
	err := DB.First(&ll, l.ID).Error
	if err != nil {
		return nil, err
	}
	err = DB.Save(&l).Error
	if err != nil {
		return nil, err
	}
	return l, err
}

func (l *Location) Remole() (err error) {
	err = DB.First(&l, l.ID).Error
	if err != nil {
		return
	}
	err = DB.Delete(&l).Error
	if err != nil {
		return
	}
	return
}
