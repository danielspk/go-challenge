package logger

import (
	"fmt"
	"go.uber.org/zap"
)

var logger *zap.Logger

// NewZapInstance crea una instancia del zap.Logger
// Nota: esta tipo de instancia singleton no funciona bien en trabajos concurrentes - Véase patrón `Singleton` en https://leanpub.com/designpatternsingo
func NewZapInstance() (*zap.Logger, error) {
	if logger != nil {
		return logger, nil
	}

	config := zap.NewProductionConfig()
	config.OutputPaths = []string{
		//"stdout",
		"./storage/logs/go_challenge.log",
	}

	newLogger, err := config.Build()

	if err != nil {
		return nil, fmt.Errorf("cannot initialize the `zap` logger: %s", err.Error())
	}

	defer newLogger.Sync()

	logger = newLogger

	return logger, nil
}
