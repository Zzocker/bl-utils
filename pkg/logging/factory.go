package logging

var (
	factory = map[string]loggerBuilder{}
)

type loggerBuilder interface {
	Build(level string, enableCaller bool) error
}

// FromFactory : returns builder of a given logger
func FromFactory(code string) loggerBuilder {
	return factory[code]
}
