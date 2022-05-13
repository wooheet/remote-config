package models

import "github.com/jinzhu/gorm"

// gen:qs
type Users struct {
	gorm.Model
	ID       uint64     `json:"id"`
	Email    string     `gorm:"uniqueIndex"`
	Password string     `json:"password"`
	UserRole []UserRole `gorm:"ForeignKey:UserId"`
}
