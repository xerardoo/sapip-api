package models

type AuditLogIncident struct {
	Model
	Action     string   `gorm:"size:500;not null;" sql:"index" json:"action"`
	UserAgent  string   `gorm:"size:500;not null;"json:"user_agent"`
	IncidentID int      `gorm:"type:integer" json:"incident_id"`
	Incident   Incident `gorm:"foreignkey:IncidentID;" json:"incident"`
	UserID     int      `gorm:"type:integer" json:"user_id"`
	User       User     `gorm:"foreignkey:UserID;" json:"user"`
	LocationID *int      `gorm:"type:integer" json:"location_id"`
	Location   Location `gorm:"foreignkey:LocationID;" json:"location"`
}

const AUDIT_INCIDENT_VISIT = "Visualización  de Incidente"

func (a *AuditLogIncident) Add() (*AuditLogIncident, error) {
	err := DB.Create(&a).Error
	if err != nil {
		return nil, err
	}
	return a, err
}
