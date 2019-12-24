package glog

import (
	"testing"
)

func TestLoadConf(t *testing.T) {
	LoadConf("examples/config.yaml")
}
