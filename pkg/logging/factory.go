package logging

import "github.com/Zzocker/bl-utils/config"

var (
	factory = map[string]loggerBuilder{
		config.ZAP: &zapBuilder{},
	}
)

type loggerBuilder interface {
	Build(cdf *LoggerConfig) error
}

// FromFactory : returns builder of a given logger
func FromFactory(code string) loggerBuilder {
	return factory[code]
}

// LoggerConfig : conatins the configuration of logger
type LoggerConfig struct {
	// Code : name of the logger
	Code          string
	Level         string
	EnabledCaller bool
	// EncodingType : json or console
	EncodingType string
	// stdout, filename
	// example ["stdout", "logger.log"]
	OutputPath []string
}
