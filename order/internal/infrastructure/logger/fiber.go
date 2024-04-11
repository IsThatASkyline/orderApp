package logger

import "go.uber.org/zap"

type FiberLogger struct {
	*Logger
}

func (l FiberLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}

func (l Logger) GetFiberLogger() FiberLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return FiberLogger{
		Logger: newSugaredLogger(logger),
	}
}
