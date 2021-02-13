package qlog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//DailyFileWriter ...
type DailyFileWriter struct {
	Path string
	Extension string
	currentFile *os.File
}

func (d *DailyFileWriter) getDailyPath() string {
	builder := strings.Builder{}
	builder.Write([]byte(d.Path))
	builder.Write([]byte(" - "))
	builder.Write([]byte(time.Now().Format("Mon, 02 Jan 2006")))
	builder.Write([]byte(d.Extension))
	return builder.String()
}

func (d *DailyFileWriter) checkFile() {
	if d.currentFile == nil {
		f, err := os.Create(d.getDailyPath())
		if err != nil {
			fmt.Println(err)
		}
		d.currentFile = f
	} else {

	}
}

func (d *DailyFileWriter) Write(p []byte) (n int, err error) {
	d.checkFile()
	return d.currentFile.Write(p)
}
