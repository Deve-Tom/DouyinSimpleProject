package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string
	UserID  uint
	VideoID uint
	User    User  `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
	Video   Video `gorm:"foreignkey:VideoID;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
}
