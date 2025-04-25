package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name    string  `gorm:"not null" json:"name" validate:"required"`
	Review  string  `json:"review,omitempty"`
	Rating  int     `json:"rating,omitempty"`
	Summary string  `json:"summary,omitempty"`
	Genre   string  `json:"genre,omitempty"`
	Actors  []Actor `gorm:"many2many:movie_actors;constraint:OnDelete:CASCADE;"`
}
