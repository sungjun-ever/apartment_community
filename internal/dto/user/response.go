package user

import (
	"apart_community/internal/model"
	"time"
)

type Resources struct {
	ID             uint      `json:"id"`
	UUID           string    `json:"uuid"`
	Email          string    `json:"email"`
	Nickname       string    `json:"nickname"`
	CreatedAt      time.Time `json:"created_at"`
	Status         int       `json:"status"`
	ProfileImageId *uint     `json:"profile_image_id"`
}

func NewResource(u *model.User) Resources {
	return Resources{
		ID:             u.ID,
		UUID:           u.UUID,
		Email:          u.Email,
		CreatedAt:      u.CreatedAt,
		Nickname:       u.Profile.Nickname,
		Status:         u.Status,
		ProfileImageId: u.Profile.ProfileImageId,
	}
}
