package models

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/vincent-petithory/dataurl"
	"gorm.io/gorm"
)

type Incident struct {
	Model
	Date          string       `gorm:"size:25;not null;" sql:"index" json:"date"`
	Description   string       `gorm:"type:text;size:500;index:,class:FULLTEXT" json:"description"` // full index narrativa
	Address       string       `gorm:"size:400;not null;" sql:"index" json:"address"`
	Area          string       `gorm:"size:25;not null;" sql:"index" json:"area"`
	ZipCode       string       `gorm:"size:25;not null;" sql:"index" json:"zipcode"`
	Type          IncidentType `gorm:"foreignkey:TypeID;" json:"type"`
	TypeID        int          `gorm:"type:integer" json:"type_id"`
	PersonasCount int          `gorm:"-" json:"personas_count"`
	VehiclesCount int          `gorm:"-" json:"vehicles_count"`
	Personas      []Persona    `gorm:"many2many:incident_personas;" json:"personas"` // involucrados
	Vehicles      []Vehicle    `gorm:"many2many:incident_vehicles;" json:"vehicles"` //
	Patrols       []Patrol     `gorm:"many2many:incident_patrols;" json:"patrols"`
	LocationID    int          `gorm:"type:integer" json:"location_id"`
	Location      Location     `gorm:"foreignkey:LocationID;" json:"location"`
	// UserID      int          `gorm:"type:integer" json:"user_id"`
	// User        User         `gorm:"foreignkey:UserID;" json:"user"`
	// fotos del incidente
}

type IncidentType struct {
	Model
	Name string `gorm:"size:250;not null;" sql:"index" json:"name"`
}

// tipo de evento  homicidio, persona lesionada, robo con violencia a comercio, robo con violencia a persona,
// robo a casa habitacion, allanamiento, rina en via publica, violencia intrafamniliar,
func InitIncidents(db *gorm.DB) {
	db.FirstOrCreate(&IncidentType{Model: Model{ID: 1}, Name: "Allamiento"})
	db.FirstOrCreate(&IncidentType{Model: Model{ID: 2}, Name: "Robo de Vehiculo"})
	db.FirstOrCreate(&IncidentType{Model: Model{ID: 3}, Name: "Robo a Casa Habitacion"})
	db.FirstOrCreate(&IncidentType{Model: Model{ID: 4}, Name: "Robo con Violencia a Comercio"})
	db.FirstOrCreate(&IncidentType{Model: Model{ID: 5}, Name: "Homicidio"})
	db.FirstOrCreate(&IncidentType{Model: Model{ID: 6}, Name: "Violencia Intrafamiliar"})
}

func (l *Incident) Add() (*Incident, error) {
	for i, p := range l.Personas {
		dataURL, err := dataurl.DecodeString(p.PhotoFront)
		if err != nil {
			return nil, err
		}
		filename := fmt.Sprintf("p_"+uuid.New().String()+".%s", dataURL.Subtype)
		err = UploadFileS3("", filename, bytes.NewReader(dataURL.Data))
		if err != nil {
			return nil, err
		}
		l.Personas[i].PhotoFront = filename
	}

	for i, v := range l.Vehicles {
		dataURL, err := dataurl.DecodeString(v.Photo)
		if err != nil {
			return nil, err
		}
		filename := fmt.Sprintf("v_"+uuid.New().String()+".%s", dataURL.Subtype)
		err = UploadFileS3("", filename, bytes.NewReader(dataURL.Data))
		if err != nil {
			return nil, err
		}
		l.Vehicles[i].Photo = filename
	}

	err := DB.Create(&l).Error
	if err != nil {
		return nil, err
	}
	return l, err
}

func (l *Incident) Find(id int) (err error) {
	err = DB.Preload("Location").Preload("Type").
		Preload("Personas.Type").Preload("Vehicles").First(&l, id).Error
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

func (l *Incident) Remove() (err error) {
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

func (l *Incident) GetLocation() (location Location, err error) {
	err = DB.Model(&l).Association("Location").Find(&location)
	if err != nil {
		return
	}
	return
}
func (l *Incident) GetUser() (user User, err error) {
	err = DB.Model(&l).Association("User").Find(&user)
	if err != nil {
		return
	}
	return
}

func (i *Incident) GetPersonasCount() (count int, err error) {
	err = DB.Raw("SELECT COUNT(*)count FROM incident_personas WHERE incident_id=?", i.ID).Scan(&count).Error
	if err != nil {
		return
	}
	return
}
func (i *Incident) GetVehiclesCount() (count int, err error) {
	err = DB.Raw("SELECT COUNT(*)count FROM incident_vehicles WHERE incident_id=?", i.ID).Scan(&count).Error
	if err != nil {
		return
	}
	return
}
