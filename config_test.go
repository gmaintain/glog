package glog

import (
	"fmt"
	"testing"
)

func TestLoadConf(t *testing.T) {
	logger, err := LoadConf("examples/config.yaml")
	fmt.Printf("%#v\n", logger)
	if err != nil {
		t.Error(err)
	}
}
