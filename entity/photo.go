package entity

import "time"

type Photo struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption  string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
