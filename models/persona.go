package models

type Persona struct {
	Model
	FirstName  string     `gorm:"size:250;not null;" sql:"index"`
	LastName   string     `gorm:"size:250;not null;" sql:"index"`
	Type       string     `gorm:"size:250;not null;" sql:"index"`
	IdentityID int        `gorm:"type:integer"`
	Identities []Persona  `gorm:"foreignkey:IdentityID"`
	// Vehicles   []Vehicle  `gorm:"many2many:persona_vehicles;"`
	// Locations  []Location `gorm:"many2many:persona_locations;"`
	// Incidents  []Incident `gorm:"many2many:persona_incidents;"`
}
