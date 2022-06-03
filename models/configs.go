package models

type Configs struct {
	//gorm.Model
	ID          uint64 `gorm:"unique" json:"id"`
	Token       string `json:"token"`
	TrackerType string `gorm:"unique" json:"tracker_type"`
	StoreId     string `gorm:"unique" json:"store_id"`
	Users       Users  `gorm:"references:ID"`
}
