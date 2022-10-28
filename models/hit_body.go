package models

type HitBody struct {
	TimeCardInfo TimeCardInfo `json:"time_card"`
	Path         string       `json:"_path"`
}

type TimeCardInfo struct {
	Accuracy          int    `json:"accuracy"`
	AccuracyMethod    bool   `json:"accuracy_method"`
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Address           string `json:"address"`
	ReferenceId       string `json:"reference_id"`
	OriginalLatitude  string `json:"original_latitude"`
	OriginalLongitude string `json:"original_longitude"`
	OriginalAddress   string `json:"original_address"`
	LocationEdited    bool   `json:"location_edited"`
}
