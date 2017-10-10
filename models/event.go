package models

import "time"

type Event struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	User string `json:"user"`
	StartDate time.Time `json:"startDate"`
	EndDate time.Time `json:"endDate"`
}