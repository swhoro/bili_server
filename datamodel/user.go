package datamodel

import (
	"gorm.io/gorm"
)

type Authority uint8

const (
	AOwner Authority = iota + 1
	AAdministrator
	AWriter
	AUser
)

type User struct {
	gorm.Model
	Username  string    `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	Authority Authority `gorm:"not null"`
}
