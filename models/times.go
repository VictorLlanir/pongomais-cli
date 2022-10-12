package models

type Times struct {
	Workday WorkdayData `json:"work_day"`
}

type WorkdayData struct {
	TimeCards []TimeCard `json:"time_cards"`
}

type TimeCard struct {
	Time    string `json:"time"`
	Date    string `json:"date"`
	Receipt string `json:"receipt"`
}
