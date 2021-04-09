package models

type Incident struct {
	Model
	Description string `gorm:"index:,class:FULLTEXT" json:"description"` // full index narrativa
	Type        string `gorm:"size:250;not null;" sql:"index" json:"type"`
	Date        string `gorm:"size:250;not null;" sql:"index" json:"date"`
	Address     string `gorm:"size:250;not null;" sql:"index" json:"address"`
	// tipo de evento  homicidio, persona lesionada, robo con violencia a comercio, robo con violencia a persona, robo a casa habitacion, allanamiento, rina en via publica, violencia intrafamniliar,
	Personas   []Persona `gorm:"many2many:incident_personas;" json:"personas"` // involucrados
	Vehicles   []Vehicle `gorm:"many2many:incident_vehicles;" json:"vehicles"` //
	Patrols    []Patrol  `gorm:"many2many:incident_patrols;" json:"patrols"`
	LocationID int       `gorm:"type:integer" json:"location_id"`
	Location   Location  `gorm:"foreignkey:LocationID;" json:"location"`
	UserID     int       `gorm:"type:integer" json:"user_id"`
	User       User      `gorm:"foreignkey:UserID;" json:"user"`
	// fotos del incidente
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
