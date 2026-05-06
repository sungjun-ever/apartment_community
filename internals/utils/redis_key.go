package utils

// RateLimitKey
/**
Client IP 또는 사용자 PublicID를 받아 레디스 키 생성
*/
func RateLimitKey(input string) string {
	return "rate_limit:" + input
}
