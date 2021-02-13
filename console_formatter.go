package qlog

import (
	"fmt"
	"io"
)

//ConsoleFormatter ...
type ConsoleFormatter struct {
	TimeFormat string
	Colors     bool
}

func levelStamp(level Level, colored bool) string {
	const inf = "[INF]"
	const infColored = (string)("[\033" + ConsoleGreen + "INF" + "\033[0m]")

	const dbg = "[DBG]"
	const dbgColored = (string)("[\033" + ConsoleYellow + "DBG" + "\033[0m]")

	const wrn = "[WRN]"
	const wrnColored = (string)("[\033" + ConsoleRed + "WRN" + "\033[0m]")

	const err = "[ERR]"
	const errColored = (string)("[\033" + ConsoleLightRed + "ERR" + "\033[0m]")

	switch level {
	case LevelInfo:
		if colored {
			return infColored
		}
		return inf

	case LevelDebug:
		if colored {
			return dbgColored
		}
		return dbg

	case LevelWarning:
		if colored {
			return wrnColored
		}
		return wrn

	case LevelError:
		if colored {
			return errColored
		}
		return err

	default:
		return ""
	}
}

//Format formats the message with a level tag, time tag, and message.
//Uses ANSI console colors if specified.
func (c ConsoleFormatter) Format(msg *Message, w io.Writer) error {
	_, err := w.Write([]byte(levelStamp(msg.Level, c.Colors)))
	if err != nil {
		return err
	}

	if c.Colors {
		const coloredTimeFmt = (string)("[\033" + ConsoleCyan + "%v" + "\033[0m]: ")
		_, err = w.Write([]byte(fmt.Sprintf(coloredTimeFmt, msg.Time.Format(c.TimeFormat))))
	} else {
		_, err = w.Write([]byte(fmt.Sprintf("[%v]: ", msg.Time.Format(c.TimeFormat))))
	}

	if err != nil {
		return err
	}
	_, err = w.Write([]byte(msg.Message))
	if err != nil {
		return err
	}

	_, err = w.Write([]byte("\n"))
	if err != nil {
		return err
	}
	return nil
}
