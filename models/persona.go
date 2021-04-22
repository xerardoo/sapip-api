package models

import (
	"gorm.io/gorm"
)

type Persona struct {
	Model
	FirstName   string      `gorm:"size:250;not null;" sql:"index" json:"first_name"`
	MiddleName  string      `gorm:"size:250;not null;" sql:"index" json:"middle_name"`
	LastName    string      `gorm:"size:250;" sql:"index" json:"last_name"`
	Alias       string      `gorm:"size:250;" sql:"index" json:"alias"`
	Sex         string      `gorm:"size:6;not null;" sql:"index" json:"sex"`
	Nationality string      `gorm:"size:6;" sql:"index" json:"nationality"`
	BirthDate   string      `gorm:"size:25;not null;" json:"birth_date"`
	Hometown    string      `gorm:"size:250;not null;" json:"hometown"`
	Occupation  string      `gorm:"size:250;not null;" json:"occupation"`
	Type        PersonaType `gorm:"foreignkey:TypeID;" json:"type"`
	TypeID      int         `gorm:"type:integer" json:"type_id"`
	// IdentityID  int         `gorm:"type:integer" json:"identity_id"`
	// Identities  []Persona   `gorm:"foreignkey:IdentityID" json:"identities"`
	// Vehicles   []Vehicle  `gorm:"many2many:persona_vehicles;"`
	// Locations  []Location `gorm:"many2many:persona_locations;"`
	// foto de frente
}

type PersonaType struct {
	Model
	Name string `gorm:"size:250;not null;" sql:"index" json:"name"`
}

// victima, reportante, fallecido, lesionado, persona responsable
func InitPersona(db *gorm.DB) {
	db.FirstOrCreate(&PersonaType{Model: Model{ID: 1}, Name: "Victima"})
	db.FirstOrCreate(&PersonaType{Model: Model{ID: 2}, Name: "Reportante"})
	db.FirstOrCreate(&PersonaType{Model: Model{ID: 3}, Name: "Lesionado"})
	db.FirstOrCreate(&PersonaType{Model: Model{ID: 4}, Name: "Responsable"})
	db.FirstOrCreate(&PersonaType{Model: Model{ID: 5}, Name: "Fallecido"})
}

func (p *Persona) Add() (*Persona, error) {
	err := DB.Create(&p).Error
	if err != nil {
		return nil, err
	}
	return p, err
}

func (p *Persona) Find(id int) (err error) {
	err = DB.First(&p, id).Error
	if err != nil {
		return
	}
	return
}

func (p *Persona) Update() (*Persona, error) {
	var pp Persona
	err := DB.First(&pp, p.ID).Error
	if err != nil {
		return nil, err
	}
	err = DB.Save(&p).Error
	if err != nil {
		return nil, err
	}
	return p, err
}

func (p *Persona) Remove() (err error) {
	err = DB.First(&p, p.ID).Error
	if err != nil {
		return
	}
	err = DB.Delete(&p).Error
	if err != nil {
		return
	}
	return
}
