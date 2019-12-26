package glog

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	logger := NewLogger()
	logger.Debug("%s", "log file 我是测试日志")
	time.Sleep(time.Second * 3)
	logger.Error("%s %s", "我是format的日志格式", ".")
	time.Sleep(time.Second * 2)
}
