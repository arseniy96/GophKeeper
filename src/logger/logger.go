package logger

import "go.uber.org/zap"

type Logger struct {
	Log *zap.SugaredLogger
}

func Initialize(l string) (*Logger, error) {
	lvl, err := zap.ParseAtomicLevel(l)
	if err != nil {
		return nil, err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return nil, err
	}
	return &Logger{Log: zl.Sugar()}, nil
}
