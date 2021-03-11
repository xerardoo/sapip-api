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
