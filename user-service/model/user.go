package model

import (
	"time"

	"gorm.io/gorm"
)

type Users []*User

type User struct {
	gorm.Model
	Name        string
	Email       string
	CreatedDate time.Time
	FirstName   string
}
