package user

import (
	"apart_community/internal/model"

	"github.com/google/uuid"
)

type UriRequest struct {
	ID uint `uri:"id" binding:"required,numeric,gt=0"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,password_check"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`

	Nickname       string `json:"nickname" binding:"required"`
	ProfileImageId *uint  `json:"profile_image_id,omitempty"`
}

func (r *RegisterRequest) ToUserEntity() model.User {
	newUUID, _ := uuid.NewV7()
	return model.User{
		UUID:     newUUID.String(),
		Email:    r.Email,
		Password: r.Password,
		Status:   1,
	}
}

func (r *RegisterRequest) ToProfileEntity() model.Profile {
	return model.Profile{
		Nickname:       r.Nickname,
		ProfileImageId: r.ProfileImageId,
	}
}
