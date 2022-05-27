package main

type Customer struct {
	Email        string `json:"email"`
	Text         string `json:"text"`
	Schedule     string `json:"schedule"`
	Paid         bool   `json:"paid"`
	ScheduleList []int  `json:"-"`
}
