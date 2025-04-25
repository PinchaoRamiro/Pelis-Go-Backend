package models

import "gorm.io/gorm"

type Serie struct {
	gorm.Model
	Name    string  `gorm:"not null" json:"name"  validate:"require"`
	Review  string  `json:"review,omitempty"`
	Rating  int     `json:"ratign,omitempty"`
	Summary string  `json:"summary,omitempty"`
	Genre   string  `json:"genre,omitempty"`
	Actors  []Actor `gorm:"many2many:series_actors;constraint:OnDelete:CASCADE;"`
}
