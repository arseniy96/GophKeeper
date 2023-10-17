package logger

import "go.uber.org/zap"

var Log *zap.SugaredLogger

func Initialize(l string) error {
	lvl, err := zap.ParseAtomicLevel(l)
	if err != nil {
		return err
	}
	cfg := zap.NewProductionConfig()
	cfg.Level = lvl
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	Log = zl.Sugar()
	return nil
}
