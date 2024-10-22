package pkg

import "go.uber.org/zap"

var Logger *zap.Logger

func InitLogger(env string) {
	baseLogger, _ := zap.NewProduction()

	Logger = baseLogger.With(zap.String("env", env))
	defer Logger.Sync()
}
