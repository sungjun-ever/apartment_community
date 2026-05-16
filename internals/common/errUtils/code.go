package errUtils

type ErrorCode string

const (
	A001 ErrorCode = "AUTH_EXPIRED_TOKEN"
	A002 ErrorCode = "AUTH_INVALID_TOKEN"
	A003 ErrorCode = "AUTH_REQUIRED"
	A004 ErrorCode = "AUTH_NO_PERMISSION"

	U001 ErrorCode = "USER_NOT_FOUND"
	U002 ErrorCode = "USER_ALREADY_EXIST"
	U003 ErrorCode = "USER_LOCKED"
	U004 ErrorCode = "USER_INVALID_PASSWORD"

	C001 ErrorCode = "INVALID_INPUT"
	C002 ErrorCode = "RESOURCE_NOT_FOUND"
	C003 ErrorCode = "CONFLICT"
	C004 ErrorCode = "TOO_MANY_REQUEST"

	S001 ErrorCode = "SYSTEM_ERROR"
	S002 ErrorCode = "SYSTEM_TIMEOUT"
	S003 ErrorCode = "SYSTEM_MAINTENANCE"
)

var responseMessage = map[ErrorCode]string{
	A001: "인증 토큰이 만료됐습니다.",
	A002: "유효하지 않은 토큰입니다.",
	A003: "로그인 정보가 없습니다.",
	A004: "권한이 없습니다.",
	U001: "사용자 정보가 없습니다,",
	U002: "이미 존재하는 사용자입니다.",
	U003: "사용자 계정이 잠겼습니다.",
	U004: "로그인 정보가 일치하지 않습니다.",
	C001: "유효하지 않은 입력입니다.",
	C002: "리소스를 찾을 수 없습니다.",
	C003: "충돌이 발생했습니다.",
	C004: "요청이 너무 많습니다.",
	S001: "시스템 에러가 발생했습니다.",
	S002: "시스템 시간 초과가 발생했습니다.",
	S003: "시스템 유지보수가 진행 중입니다.",
}

func (e ErrorCode) String() string {
	return string(e)
}

func (e ErrorCode) GetMessage() string {
	if msg, ok := responseMessage[e]; ok {
		return msg
	}

	return "UNKNOWN_ERROR"
}
