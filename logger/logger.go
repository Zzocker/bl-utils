package logger

type Level int

const (
	LEVEL_DEBUG Level = -1
	LEVEL_INFO
	LEVEL_ERROR
)

type Interface interface {
	Debug(...string)
	Debugf(string, ...string)
	Info(...string)
	Infof(string, ...string)
	Error(...string)
	Errorf(string, ...string)
}
