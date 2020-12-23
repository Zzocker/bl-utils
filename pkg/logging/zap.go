package logging

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type zapBuilder struct{}

func (z *zapBuilder) Build(level string, enableCaller bool) error {
	zLogger, err := createZAP(level, enableCaller)
	if err != nil {
		return err
	}
	defer zLogger.Sync()
	zSugarlog := zLogger.Sugar()
	zSugarlog.Info()

	//This is for loggerWrapper implementation
	//appLogger.SetLogger(&loggerWrapper{zaplog})

	setLogger(zSugarlog)
	return nil
}

func createZAP(level string, enabledCaller bool) (zap.Logger, error) {

	rawJSON := []byte(`{
		"level": "info",
		"Development": true,
		"DisableCaller": false,
		"encoding": "console",
		"outputPaths": ["stdout", "../../../demo.log"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		   "timeKey":        "ts",
		   "levelKey":       "level",
		   "messageKey":     "msg",
			"nameKey":        "name",
		   "stacktraceKey":  "stacktrace",
			"callerKey":      "caller",
		   "lineEnding":     "\n\t",
		   "timeEncoder":     "time",
		   "levelEncoder":    "lowercaseLevel",
		   "durationEncoder": "stringDuration",
			"callerEncoder":   "shortCaller"
		}
	   }`)

	var cfg zap.Config
	var zLogger *zap.Logger
	//standard configuration
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return *zLogger, errors.Wrap(err, "Unmarshal")
	}
	//customize it from configuration file
	err := customizeLogFromConfig(&cfg, level, enabledCaller)
	if err != nil {
		return *zLogger, err
	}
	zLogger, err = cfg.Build()
	if err != nil {
		return *zLogger, err
	}

	zLogger.Debug("logger construction succeeded")
	return *zLogger, nil
}

func customizeLogFromConfig(cfg *zap.Config, level string, enabledCaller bool) error {
	cfg.DisableCaller = !enabledCaller

	// set log level
	l := zap.NewAtomicLevel().Level()
	err := l.Set(level)
	if err != nil {
		return err
	}
	cfg.Level.SetLevel(l)
	return nil
}
