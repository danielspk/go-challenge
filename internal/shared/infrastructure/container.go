package infrastructure

import (
	"challenge.com/challenge/internal/search/infrastructure/graphhopper"
	"challenge.com/challenge/internal/shared/infrastructure/persistence/mysql"
	"go.uber.org/zap"
)

// Container estructura de dependencias
type Container struct {
	Logger      *zap.Logger
	Database    *mysql.Client
	GraphHopper *graphhopper.Service
}
