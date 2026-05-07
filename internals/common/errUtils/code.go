package errUtils

var (
	// A001 expired token
	A001 = "AUTH_EXPIRED_TOKEN"
	// A002 invalid token
	A002 = "AUTH_INVALID_TOKEN"
	// A003 auth required
	A003 = "AUTH_REQUIRED"
	// A004 no permission
	A004 = "AUTH_NO_PERMISSION"

	// U001 user not found
	U001 = "USER_NOT_FOUND"
	// U002 user already exist
	U002 = "USER_ALREADY_EXIST"
	// U003 user locked
	U003 = "USER_LOCKED"
	// U004 user invalid password
	U004 = "USER_INVALID_PASSWORD"

	// C001 invalid input
	C001 = "INVALID_INPUT"
	// C002 resource not found
	C002 = "RESOURCE_NOT_FOUND"
	// C003 conflict
	C003 = "CONFLICT"
	// C004 too many request
	C004 = "TOO_MANY_REQUEST"

	// S001 system error
	S001 = "SYSTEM_ERROR"
	// S002 system timeout
	S002 = "SYSTEM_TIMEOUT"
	// S003 system maintenance
	S003 = "SYSTEM_MAINTENANCE"
)

var errorMessage = map[string]string{
	A001: "인증 토큰이 만료됐습니다.",
	A002: "유효하지 않은 토큰입니다.",
	A003: "로그인 정보가 없습니다.",
	A004: "권한이 없습니다.",
	U001: "사용자 정보가 없습니다,",
	U002: "이미 존재하는 사용자입니다.",
	U003: "사용자 계정이 잠겼습니다.",
	U004: "비밀번호가 일치하지 않습니다.",
	C001: "유효하지 않은 입력입니다.",
	C002: "리소스를 찾을 수 없습니다.",
	C003: "충돌이 발생했습니다.",
	C004: "요청이 너무 많습니다.",
	S001: "시스템 에러가 발생했습니다.",
	S002: "시스템 시간 초과가 발생했습니다.",
	S003: "시스템 유지보수가 진행 중입니다.",
}

func GetErrorMessage(code string) string {
	if msg, ok := errorMessage[code]; ok {
		return msg
	}

	return "알 수 없는 에러입니다."
}
