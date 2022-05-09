package models

import "github.com/jinzhu/gorm"

// gen:qs
type Users struct {
	gorm.Model
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"password"`
	UserRole string `json:"user_role"`
}

type UserRole struct {
	ID     uint64 `json:"id" gorm:"primaryKey"`
	UserId uint64 `json:"user_id"`
	Role   string `json:"role"`
}
