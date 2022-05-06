package models

import "github.com/jinzhu/gorm"

// gen:qs
type Users struct {
	gorm.Model
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"primaryKey"`
	Password string `json:"password"`
}
