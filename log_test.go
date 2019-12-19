package glog

import (
	"testing"
)

func TestLogger(t *testing.T)  {
	logger := NewLogger()
	logger.Debug("%s", "log file")
	//time.Sleep(time.Second * 1)
}