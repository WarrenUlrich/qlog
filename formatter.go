package qlog

import "io"

//Formatter interface for formatting and writing a log message
//to an io.Writer.
type Formatter interface {
	Format(msg *Message, w io.Writer) error
}