package logger

type Logger struct {
	Stdout bool   `json:"stdout" xml:"stdout" yaml:"stdout" ini:"STDOUT" comment:"是否输出到控制台"`
	Level  int    `json:"level" xml:"level" yaml:"level" ini:"LEVEL" comment:"日志级别"`
	File   string `json:"file" xml:"file" yaml:"file" ini:"FILE" comment:"日志文件"`
}
