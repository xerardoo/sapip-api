package models

type Persona struct {
	Model
	FirstName  string    `gorm:"size:250;not null;" sql:"index" json:"first_name"`
	LastName   string    `gorm:"size:250;not null;" sql:"index" json:"last_name"`
	Type       string    `gorm:"size:250;not null;" sql:"index" json:"type"`
	IdentityID int       `gorm:"type:integer" json:"identity_id"`
	Identities []Persona `gorm:"foreignkey:IdentityID" json:"identities"`
	// Vehicles   []Vehicle  `gorm:"many2many:persona_vehicles;"`
	// Locations  []Location `gorm:"many2many:persona_locations;"`
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

func (p *Persona) Remole() (err error) {
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
