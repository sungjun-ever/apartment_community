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

type BelongApartmentRequest struct {
	UserID      uint `json:"user_id" binding:"required"`
	ApartmentID uint `json:"apartment_id" binding:"required"`
	RoleID      uint `json:"role_id" binding:"required"`
	Unit        uint `json:"unit" bindind:"omitempty"`
	No          uint `json:"no" bindind:"omitempty"`
	IsVerified  bool `json:"is_verified" binding:"omitempty"`
}

func (r *RegisterRequest) ToUserEntity() *model.User {
	newUUID, _ := uuid.NewV7()
	return &model.User{
		UUID:     newUUID.String(),
		Email:    r.Email,
		Password: r.Password,
		Status:   1,
	}
}

func (r *RegisterRequest) ToProfileEntity() *model.Profile {
	return &model.Profile{
		Nickname:       r.Nickname,
		ProfileImageId: r.ProfileImageId,
	}
}

func (r *BelongApartmentRequest) ToUbaEntity() *model.UserBelongApartment {
	return &model.UserBelongApartment{
		UserID:      r.UserID,
		ApartmentID: r.ApartmentID,
		RoleID:      r.RoleID,
		Unit:        &r.Unit,
		No:          &r.No,
		IsVerified:  r.IsVerified,
	}
}
