package qlog

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//DailyFileWriter is an io.Writer implementation
//that creates a file with the provided path and extension,
//and an added date, e.g. file_name - Mon, 02 Jan 2006.txt.
//Creates a new file if the day changes.
type DailyFileWriter struct {
	Path string
	currentFile *os.File
	updateTime time.Time
}

func (d *DailyFileWriter) getDailyPath() string {
	builder := strings.Builder{}

	seperated := strings.Split(d.Path, ".")

	builder.Write([]byte(seperated[0]))
	builder.Write([]byte(" - "))
	builder.Write([]byte(time.Now().Format("Mon, 02 Jan 2006.")))
	builder.Write([]byte(seperated[1]))
	return builder.String()
}

func (d *DailyFileWriter) checkFile() {
	if d.currentFile == nil {
		f, err := os.Create(d.getDailyPath())
		if err != nil {
			fmt.Println(err)
		}
		d.currentFile = f
		now := time.Now().Round(time.Hour)
		d.updateTime = now.Add((24 - (time.Duration)(now.Hour())) * time.Hour)
	} else {
		if time.Now().After(d.updateTime) {
			d.currentFile.Close()
			d.currentFile = nil
			d.checkFile()
		}
	}
}

func (d *DailyFileWriter) Write(p []byte) (n int, err error) {
	d.checkFile()
	return d.currentFile.Write(p)
}
