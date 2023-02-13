package entity

import "gorm.io/gorm"

// Video has a foreign key to User (one to many)
type Video struct {
	gorm.Model
	UserID        uint
	User          User `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
	Title         string
	PlayURL       string `gorm:"size:255;not null"`
	CoverURL      string `gorm:"size:255;not null"`
	FavoriteCount uint   `gorm:"default:0"`
	CommentCount  uint   `gorm:"default:0"`
}
