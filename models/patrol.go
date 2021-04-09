package models

type Patrol struct {
	Model
	UnitNumber string `gorm:"size:32" json:"unit_number"`
	IsActive   int    `gorm:"type:integer" json:"is_active"`
	Cops       []User `gorm:"many2many:patrol_user;" json:"cops"`
}
