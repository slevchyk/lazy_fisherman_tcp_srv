package models

import "time"

type MapMarker struct {
	Id        string    `json:"id"`
	AppUserId string    `json:"app_user_id"`
	ExtId     string    `json:"ext_id"`
	BoardId   string    `json:"board_id"`
	WaterId   string    `json:"water_id"`
	Type      int       `json:"type"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Title     string    `json:"title"`
	Snippet   string    `json:"snippet"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}