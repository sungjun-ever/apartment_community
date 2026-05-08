package domain

type PublicIdUriRequest struct {
	PublicID string `uri:"publicID"`
}

type PaginationRequest struct {
	Page   int  `form:"page,default=1" binding:"min=1"`
	Size   int  `form:"size,default=10" binding:"min=1,max=100"`
	IsDesc bool `form:"isDesc;default=true"`
}

type UserAuthBase struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=16"`
}

type ProfileBase struct {
	NickName string `json:"nickname" binding:"omitempty,min=2,max=16"`
	ImageID  *uint  `json:"imageId" binding:"omitempty"`
}

type RegisterRequest struct {
	UserAuthBase
	ProfileBase
	PasswordConfirm string `json:"passwordConfirm" binding:"required,eqfield=Password"`
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
