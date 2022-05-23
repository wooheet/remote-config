package models

type Configs struct {
	//gorm.Model
	ID          uint64 `gorm:"unique" json:"id"`
	Token       string `json:"token"`
	TrackerType string `json:"tracker_type"`
	StoreId     string `json:"store_id"`
}
