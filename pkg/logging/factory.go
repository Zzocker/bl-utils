package logging

import "github.com/Zzocker/bl-utils/config"

var (
	factory = map[string]loggerBuilder{
		config.ZAP: &zapBuilder{},
	}
)

type loggerBuilder interface {
	Build(level string, enableCaller bool) error
}

// FromFactory : returns builder of a given logger
func FromFactory(code string) loggerBuilder {
	return factory[code]
}
