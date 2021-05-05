package logger

import (
	"fmt"
	"io"
	"log"
	"strings"
)

const (
	goLog_callDepth = 2
)

// uses log package
type goLog struct {
	level Level
	// debug logger
	dlog *log.Logger
	// info logger
	ilog *log.Logger
	// error logger
	elog *log.Logger
}

// NewGoLog : returns logger which uses go built in log package
func NewGoLog(level Level, out io.Writer, err io.Writer) Interface {
	glog := goLog{
		level: level,
		dlog:  log.New(out, "[DEBUG] ", log.Lshortfile),
		ilog:  log.New(out, "[INFO] ", log.Lshortfile),
		elog:  log.New(err, "[ERROR] ", log.Lshortfile),
	}
	return &glog
}

func (l *goLog) Debug(v ...string) {
	if l.level > LEVEL_DEBUG {
		return
	}
	_ = l.dlog.Output(goLog_callDepth, strings.Join(v, " "))
}
func (l *goLog) Debugf(f string, v ...string) {
	if l.level > LEVEL_DEBUG {
		return
	}
	_ = l.dlog.Output(goLog_callDepth, fmt.Sprintf(f, v))
}
func (l *goLog) Info(v ...string) {
	if l.level > LEVEL_INFO {
		return
	}
	_ = l.ilog.Output(goLog_callDepth, strings.Join(v, " "))
}
func (l *goLog) Infof(f string, v ...string) {
	if l.level > LEVEL_INFO {
		return
	}
	_ = l.ilog.Output(goLog_callDepth, fmt.Sprintf(f, v))
}
func (l *goLog) Error(v ...string) {
	if l.level > LEVEL_ERROR {
		return
	}
	_ = l.elog.Output(goLog_callDepth, strings.Join(v, " "))
}
func (l *goLog) Errorf(f string, v ...string) {
	if l.level > LEVEL_ERROR {
		return
	}
	_ = l.elog.Output(goLog_callDepth, fmt.Sprintf(f, v))
}
