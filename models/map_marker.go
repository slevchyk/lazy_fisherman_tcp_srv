package models

import "time"

type MapMarker struct {
	Id        int64       `json:"id"`	
	AppUserId int64       `json:"app_user_id"`	
	BoardId   int64       `json:"board_id"`
	WaterId   int64       `json:"water_id"`
	Type      int64       `json:"type"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	Title     string    `json:"title"`
	Snippet   string    `json:"snippet"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
