package models

import "github.com/jinzhu/gorm"

type Configs struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	Token       string `json:"token"`
	TrackerType string `json:"tracker_type"`
	StoreId     string `json:"store_id"`
}
