package glog

import (
	"bufio"
	"fmt"
)

type FileWriter struct {
	filebufWriter *bufio.Writer
}

func (f *FileWriter) Flush() error {
	if f.filebufWriter != nil {
		return f.filebufWriter.Flush()
	}
	return nil
}

func (f *FileWriter) Init() error {
	panic("implement me")
}

func (f *FileWriter) Write(r *Record) error {
	fmt.Println("file write", "write func")
	if _, err := f.filebufWriter.WriteString(r.String()); err != nil {
		return err
	}
	return nil
}

func NewFileWriter() *FileWriter {
	return &FileWriter{}
}
