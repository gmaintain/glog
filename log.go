package glog

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var LEVEL_FLAGS = [...]string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR", "FATAL"}

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
)

type Record struct {
	time string
	msg string
	info string
	level int
}

type Logger struct {
	level int
	recordPool  *sync.Pool
	tunnel chan *Record
}

func NewLogger() *Logger {
	l := new(Logger)
	l.level = DEBUG
	return l
}

func (l *Logger) Debug(fmt string, args ...interface{})  {
	l.deliveryRecordTowriter(DEBUG, fmt, args)
}

func (l *Logger) deliveryRecordTowriter(level int, format string, args ...interface{})  {
	var msg string

	info := fmt.Sprintf(format, args)
	_, file, line, ok := runtime.Caller(3)
	if ok {
		msg = path.Base(file) + ":" + strconv.Itoa(line)
	}
	now := time.Now()
	r := l.recordPool.Get().(*Record)
	r.msg = msg
	r.info = info
	r.level = level
	r.time = now
	l.tunnel <- r
}