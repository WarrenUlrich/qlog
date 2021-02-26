package qlog

//Level represents log severity
type Level uint8

const (
	//LevelInfo represents info level severity.
	LevelInfo Level = iota

	//LevelDebug represents debug level severity.
	LevelDebug Level = iota

	//LevelWarning represents warning level severity.
	LevelWarning Level = iota
	
	//LevelError represents error level severity.
	LevelError Level = iota
)

func (l Level) String() string {
	switch l {
	case LevelInfo:
		return "info"

	case LevelDebug:
		return "debug"

	case LevelWarning:
		return "warning"

	case LevelError:
		return "error"

	default:
		return ""
	}
}