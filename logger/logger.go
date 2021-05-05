package logger

type Level int

const (
	LEVEL_DEBUG Level = -1
	LEVEL_INFO
	LEVEL_ERROR
)

type Interface interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}
