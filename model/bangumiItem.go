package model

import "gorm.io/gorm"

type BangumiItem struct {
	gorm.Model
	UserID   uint
	CretedBy User   `gorm:"foreignKey:UserID;references:ID"`
	Name     string `gorm:"not null;unique"`
	WebUrl   string `gorm:"not null"`
	PicUrl   string `gorm:"not null"`
}
