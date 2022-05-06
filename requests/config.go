package requests

type Config struct {
	Token       string `json:"token"`
	TrackerType string `json:"tracker_type"`
	StoreId     string `json:"store_id"`
}
