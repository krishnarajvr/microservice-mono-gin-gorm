package model

import (
	"time"

	"gorm.io/gorm"
)

type Token struct {
	gorm.Model
	TenantId  uint
	Name      string
	Email     string
	FirstName string
	LastName  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
