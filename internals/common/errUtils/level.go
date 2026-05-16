package errUtils

type ErrorLevel string

const (
	LevelDebug ErrorLevel = "debug"
	LevelInfo  ErrorLevel = "info"
	LevelWarn  ErrorLevel = "warn"
	LevelFatal ErrorLevel = "fatal"
)

func (e ErrorLevel) String() string {
	return string(e)
}
