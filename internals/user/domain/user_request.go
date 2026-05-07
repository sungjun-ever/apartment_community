package domain

type ProfileRequest struct {
	NickName string `json:"name"`
	ImageID  *uint  `json:"imageId" binding:"omitempty"`
}
type UserRequest struct {
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6,max=16"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required,eqfield=Password"`
}

type RegisterRequest struct {
	UserRequest
	ProfileRequest
}

func (r *RegisterRequest) ToUserEntity() *User {
	return &User{
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r *RegisterRequest) ToProfileEntity() *Profile {
	return &Profile{
		Nickname: r.NickName,
		ImageID:  r.ImageID,
	}
}
