package user

import (
	"apart_community/internal/model"

	"github.com/google/uuid"
)

type UriRequest struct {
	ID uint `uri:"id" binding:"required,numeric,gt=0"`
}

type ProfileRequest struct {
	Nickname       string `json:"nickname" binding:"required"`
	ProfileImageId *uint  `json:"profileImageId" binding:"omitempty"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,password_check"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`

	ProfileRequest
}

type UpdatePasswordRequest struct {
	OriginPassword     string `json:"originPassword" binding:"required,password_check"`
	NewPassword        string `json:"newPassword" binding:"required,password_check"`
	NewPasswordConfirm string `json:"newPasswordConfirm" binding:"required,eqfield=NewPassword"`
}

type BelongApartmentRequest struct {
	UserID      uint `json:"userId" binding:"required"`
	ApartmentID uint `json:"apartmentId" binding:"required"`
	RoleID      uint `json:"roleId" binding:"required"`
	Unit        uint `json:"unit" binding:"omitempty"`
	No          uint `json:"no" binding:"omitempty"`
	IsVerified  bool `json:"isVerified" binding:"omitempty"`
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

func (r *ProfileRequest) ToProfileEntity() *model.Profile {
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
