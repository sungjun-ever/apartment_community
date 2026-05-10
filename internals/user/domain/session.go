package domain

type UserSession struct {
	RefreshToken string `json:"refreshToken"`
	IP           string `json:"ip"`
	UserAgent    string `json:"userAgent"`
	CreatedAt    string `json:"createdAt"`
}
