package glog

type WriterConfig struct {
	Enable bool
	Type string
	FileName string
	FilePath string
	MaxLines int64
	MaxSize string
	Level string
}
