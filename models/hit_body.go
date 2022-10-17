package models

type HitBody struct {
	TimeCardInfo TimeCardInfo `json:"time_card"`
	Path         string       `json:"_path"`
	Device       DeviceInfo   `json:"_device"`
	AppVersion   string       `json:"_appVersion"`
}

type TimeCardInfo struct {
	Latitude          string `json:"latitude"`
	Longitude         string `json:"longitude"`
	Address           string `json:"address"`
	ReferenceId       string `json:"reference_id"`
	OriginalLatitude  string `json:"original_latitude"`
	OriginalLongitude string `json:"original_longitude"`
	OriginalAddress   string `json:"original_address"`
	LocationEdited    bool   `json:"location_edited"`
}

type DeviceInfo struct {
	Browser BrowserInfo `json:"browser"`
}

type BrowserInfo struct {
	Name                string `json:"name"`
	Version             string `json:"version"`
	VersionSearchString string `json:"versionSearchString"`
}
