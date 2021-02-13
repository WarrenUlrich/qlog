package qlog

import "io"

//Formatter ...
type Formatter interface {
	Format(msg *Message, w io.Writer) error
}