package entity

import "time"

type SocialMedias struct {
	Id        int       `json:"id"`
	Name     string  `json:"name"`
	SocialMediaUrl  string   `json:"social_media_url"`
	UserId  int    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
