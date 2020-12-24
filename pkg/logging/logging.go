package logging

import "time"

// Log is a package level variable
var Log logger

type logger interface {
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

func setLogger(log logger) {
	Log = log
}

// Profilers

// Duration : Timing profiling
func Duration(invocation time.Time, caller string) {
	Log.Debugf("%s took %d ns", caller, time.Since(invocation).Nanoseconds())
}
