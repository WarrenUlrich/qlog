package qlog

import "time"

//Message ...
type Message struct {
	Level Level
	Time time.Time
	Message string
}

