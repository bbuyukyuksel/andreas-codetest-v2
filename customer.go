package main

type Customer struct {
	Email        string `json:"email"`
	Text         string `json:"text"`
	Schedule     string `json:"-"`
	Paid         bool   `json:"paid"`
	ScheduleList []int  `json:"-"`
}
