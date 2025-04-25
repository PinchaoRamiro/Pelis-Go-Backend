package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"not null" json:"name" validate:"required"`
	Password      string `gorm:"not null" json:"-"`
	Email         string `gorm:"unique;not null" json:"email" validate:"required,email"`
	Genre         string `json:"genre,omitempty"`
	TypeOfContent string `json:"typeOfContent,omitempty"`
	Banned        bool   `json:"banned,omitempty" gorm:"default:false"`
}

type Favorite struct {
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null;index"`
	ContentID   uint   `gorm:"not null"`
	ContentType string `gorm:"not null"` // e.g., "actor", "movie", "serie"
	CreatedAt   time.Time
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
