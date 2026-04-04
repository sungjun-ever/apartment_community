package user

import "apart_community/internal/model"

type UriRequest struct {
	ID uint `uri:"id" binding:"required,numeric,gt=0"`
}

type RegisterRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,password_check"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

func (r *RegisterRequest) ToEntity() model.User {
	return model.User{
		Email:    r.Email,
		Password: r.Password,
	}
}
