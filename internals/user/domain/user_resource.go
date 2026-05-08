package domain

import "time"

type UserResource struct {
	PublicID  string `json:"pid"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	ImageID   *uint  `json:"imageId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func NewUserResource(u *User) *UserResource {
	return &UserResource{
		PublicID:  u.PublicID,
		Email:     u.Email,
		Nickname:  u.Profile.Nickname,
		ImageID:   u.Profile.ImageID,
		CreatedAt: u.CreatedAt.Format(time.DateTime),
		UpdatedAt: u.UpdatedAt.Format(time.DateTime),
	}
}

func NewUserResources(users []*User) []*UserResource {
	resources := make([]*UserResource, 0, len(users))
	for _, user := range users {
		resources = append(resources, NewUserResource(user))
	}
	return resources
}
