package models

type Config struct {
	ID          uint64 `json:"id"`
	Token       string `json:"token"`
	TrackerType string `json:"tracker_type"`
	StoreId     string `json:"store_id"`
	AtCreated   int64  `json:"at_created"`
	AtUpdated   int64  `json:"at_updated"`
}
