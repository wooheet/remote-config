package requests

type Configs struct {
	Token       string `json:"token"`
	TrackerType string `json:"tracker_type"`
	StoreId     string `json:"store_id"`
}

type ScriptTag struct {
	Script string `json:"script"`
}
