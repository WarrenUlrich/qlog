package qlog

import (
	"os"
)

//FileWriter io.Writer implementation for
//writing to a file with a specified path and extension.
type FileWriter struct {
	Path      string
	file      *os.File
}

func (f *FileWriter) checkFile() error {
	if f.file == nil {
		file, err := os.Create(f.Path)
		if err != nil {
			return err
		}
		f.file = file
	}

	return nil
}

func (f *FileWriter) Write(p []byte) (n int, err error) {
	f.checkFile()
	return f.file.Write(p)
}
