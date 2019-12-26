package glog

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Printf("file write record: %v\n", r.String())
	if _, err := f.filebufWriter.WriteString(r.String()); err != nil {
		return err
	}
	return nil
}

func NewFileWriter() *FileWriter {
	file, err := os.OpenFile("out/log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	return &FileWriter{bufio.NewWriter(file)}
}
