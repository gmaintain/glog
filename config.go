package glog

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type LogConfig []WriterConfig

type WriterConfig struct {
	Enable   bool   `yaml:"enable"`
	Type     string `yaml:"type"`
	FileName string `yaml:"file_name"`
	FilePath string `yaml:"file_path"`
	MaxLines int64  `yaml:"max_lines"`
	MaxSize  string `yaml:"max_size"`
	Level    string `yaml:"level"`
	Color    bool   `yaml:"color"` // 用在console中,true代表带颜色输出
}

func LoadConf(confPath string) (*Logger, error) {
	if _, err := os.Stat(confPath); err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}
	var confs LogConfig

	err = yaml.Unmarshal(data, &confs)
	if err != nil {
		return nil, err
	}
	l := NewLogger()

	for _, conf := range confs {
		if conf.Enable == false {
			continue
		}
		l.level = LEVEL_MAP[conf.Level]
		switch conf.Type {
		case "file":
			l.writers = append(l.writers, NewFileWriter())
		case "console":
			//todo
		}
	}
	return l, nil
}
