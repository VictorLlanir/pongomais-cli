package models

type HitBody struct {
	TimeCardInfo TimeCard `json:"time_card"`
	TimeCardInfo TimeCard `json:"time_card"`
}

type TimeCardInfo struct {
	Latitude    string `json:"latitude`
	Longitude   string `json:"longitude`
	Address     string `json:"address"`
	ReferenceId string `json:"reference_id"`
}
