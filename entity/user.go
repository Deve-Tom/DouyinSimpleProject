package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string    `gorm:"size:32;not null;unique"`
	Password        string    `gorm:"size:200;not null"`
	Nickname        string    `gorm:"size:50;not null"`
	FollowCount     uint      `gorm:"default:0"`
	FollowerCount   uint      `gorm:"default:0"`
	WorkCount       uint      `gorm:"default:0"`
	FavoriteCount   uint      `gorm:"default:0"`
	TotalFavorCount uint      `gorm:"default:0"`
	Videos          []Video   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
	FavoriteVideos  []Video   `gorm:"many2many:favorite;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
	Comments        []Comment `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
	FollowUsers     []*User   `gorm:"many2many:follow;constraint:onUpdate:CASCADE,onDELETE:CASCADE"`
}
