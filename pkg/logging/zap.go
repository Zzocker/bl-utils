package logging

import (
	"encoding/json"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type zapBuilder struct{}

func (z *zapBuilder) Build(cfg *LoggerConfig) error {
	zLogger, err := createZAP(cfg)
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

func createZAP(lCfg *LoggerConfig) (zap.Logger, error) {

	rawJSON := []byte(`{
		"level": "info",
		"Development": true,
		"DisableCaller": false,
		"encoding": "console",
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"],
		"encoderConfig": {
		   "timeKey":        "ts",
		   "levelKey":       "level",
		   "messageKey":     "msg",
			"nameKey":        "name",
		   "stacktraceKey":  "stacktrace",
			"callerKey":      "caller",
		   "lineEnding":     "\n",
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
	err := customizeLogFromConfig(&cfg, lCfg)
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

func customizeLogFromConfig(cfg *zap.Config, lCfg *LoggerConfig) error {
	cfg.DisableCaller = !lCfg.EnabledCaller

	// set log level
	l := zap.NewAtomicLevel().Level()
	err := l.Set(lCfg.Level)
	if err != nil {
		return err
	}
	cfg.Level.SetLevel(l)
	cfg.Encoding = lCfg.EncodingType
	cfg.OutputPaths = lCfg.OutputPath
	return nil
}
