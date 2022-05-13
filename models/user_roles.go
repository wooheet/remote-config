package models

import "github.com/jinzhu/gorm"

type UserRole struct {
	gorm.Model
	ID     uint64 `json:"id"`
	UserId uint64 `json:"user_id"`
	Role   string `json:"role"`
}
