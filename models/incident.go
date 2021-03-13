package models

type Incident struct {
	Model
	Description string    `gorm:"index:,class:FULLTEXT"` // full index
	Type        string    `gorm:"size:250;not null;" sql:"index"`
	Personas    []Persona `gorm:"many2many:incident_personas;"`
	Vehicles    []Vehicle `gorm:"many2many:incident_vehicles;"`
	LocationID  int       `gorm:"type:integer"`
	Location    Location  `gorm:"foreignKey:LocationID"`
	UserID      int       `gorm:"type:integer"`
	User        User      `gorm:"foreignKey:UserID"`
}

func (l *Incident) Add() (*Incident, error) {
	err := DB.Create(&l).Error
	if err != nil {
		return nil, err
	}
	return l, err
}

func (l *Incident) Find(id int) (err error) {
	err = DB.First(&l, id).Error
	if err != nil {
		return
	}
	return
}

func (l *Incident) Update() (*Incident, error) {
	var ll Incident
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

func (l *Incident) Remole() (err error) {
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
