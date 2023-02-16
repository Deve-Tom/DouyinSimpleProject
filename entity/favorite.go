package entity

import "gorm.io/gorm"

// Favorite is a customized JoinTable.
//
// GROM will automatically generate a JoinTable according to `many2many`,
// but it is simple, just has two fields user_id and video_id.
// Setup it using `SetupJoinTable` in config.go.
type Favorite struct {
	gorm.Model
	UserID  uint `gorm:"primaryKey"`
	VideoID uint `gorm:"primaryKey"`
}
