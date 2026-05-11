package utils

// RateLimitKey
/**
Client IP 또는 사용자 PublicID를 받아 레디스 키 생성
*/
func RateLimitKey(input string) string {
	return "rate_limit:" + input
}

// 사용자 인증 관련 레디스 키
func SessionKey(publicID string) string {
	return "auth:refresh:" + publicID
}
func BlacklistAccessTokenKey(accessToken string) string {
	return "auth:blacklist:" + accessToken
}
