package glog

import (
	"fmt"
	"log"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Writer interface {
	Write(*Record) error
}

type Flusher interface {
	Flush() error
}

type Rotated interface {
	Rotate() error
	SetLocation(string) error
}

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

var LEVEL_STRING = [...]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

var LEVEL_MAP = map[string]int{
	"TRACE": 1,
	"DEBUG": 2,
	"INFO":  3,
	"WARN":  4,
	"ERROR": 5,
	"FATAL": 6,
}

type Record struct {
	timeinfo string
	msg      string
	info     string
	level    int
}

func (r *Record) String() string {
	return fmt.Sprintf("[%s] [%s] [%s] [%s]", LEVEL_STRING[r.level], r.timeinfo, r.info, r.msg)
}

type Logger struct {
	level      int
	recordPool *sync.Pool
	c          chan bool
	tunnel     chan *Record
	timeLayout string
	writers    []Writer
}

func NewLogger() *Logger {
	l := new(Logger)
	l.timeLayout = "2016-01-02 15:04:05"
	l.recordPool = &sync.Pool{New: func() interface{} {
		return &Record{}
	}}
	l.tunnel = make(chan *Record)
	l.writers = append(l.writers, NewFileWriter())
	go writerRunner(l)
	return l
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.dispatch(DEBUG, format, args...)
}

func (l *Logger) Trace(format string, args ...interface{}) {
	l.dispatch(TRACE, format, args...)
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.dispatch(INFO, format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.dispatch(WARN, format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.dispatch(ERROR, format, args...)
}

func (l *Logger) Fatal(format string, args ...interface{}) {
	l.dispatch(FATAL, format, args...)
}

func (l *Logger) dispatch(level int, format string, args ...interface{}) {
	var msg string

	info := fmt.Sprintf(format, args...)
	_, file, line, ok := runtime.Caller(3)
	if ok {
		msg = path.Base(file) + ":" + strconv.Itoa(line)
	}
	now := time.Now()
	r := l.recordPool.Get().(*Record)

	r.timeinfo = now.Format(l.timeLayout)
	r.msg = msg
	r.info = info
	r.level = level

	l.tunnel <- r
}

func writerRunner(logger *Logger) {
	if logger == nil {
		panic("logger is nil")
	}
	flushtimer := time.NewTimer(time.Millisecond * 500)
	rotatetimer := time.NewTimer(time.Hour * 24)

	for {
		select {
		case record, ok := <-logger.tunnel:
			fmt.Println("ok :", ok)
			if !ok {
				logger.c <- true
				return
			}
			for _, w := range logger.writers {
				fmt.Println(w)
				if err := w.Write(record); err != nil {
					log.Println(err)
				}
			}
		case <-flushtimer.C:
			for _, w := range logger.writers {
				if f, ok := w.(Flusher); ok {
					if err := f.Flush(); err != nil {
						log.Println("flush error:", err)
					}
				}
			}
		case <-rotatetimer.C:
			for _, w := range logger.writers {
				if f, ok := w.(Rotated); ok {
					if err := f.Rotate(); err != nil {
						log.Println("log rotate:", err)
					}
				}
			}
		}

	}
}
