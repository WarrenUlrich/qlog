package qlog

import (
	"os"
)

//FileWriter ...
type FileWriter struct {
	Path      string
	Extension string
	file      *os.File
}

func (f *FileWriter) checkFile() error {
	if f.file == nil {
		file, err := os.Create(f.Path + f.Extension)
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
