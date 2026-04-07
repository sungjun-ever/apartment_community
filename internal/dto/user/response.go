package user

import "time"

type Resources struct {
	ID             uint      `json:"id"`
	UUID           string    `json:"uuid"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	CreatedAt      time.Time `json:"created_at"`
	Status         int       `json:"status"`
	ProfileImageId *uint     `json:"profile_image_id"`
}
