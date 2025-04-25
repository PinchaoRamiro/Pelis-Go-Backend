package models

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	Name   string  `gorm:"not null" json:"name" validate:"required"`
	Rating int     `json:"rating,omitempty"`
	About  string  `json:"about,omitempty"`
	Movies []Movie `gorm:"many2many:movie_actors;constraint:OnDelete:CASCADE;"`
	Series []Serie `gorm:"many2many:series_actors;constraint:OnDelete:CASCADE;"`
}
