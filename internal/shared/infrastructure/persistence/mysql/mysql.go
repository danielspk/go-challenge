package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// Client estructura de cliente de conexión a MySQL
type Client struct {
	*sql.DB
}

// NewMysqlClient crea un Client
func NewMysqlClient(dataSourceName string) (*Client, error) {
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		return nil, fmt.Errorf("cannot create `mysql` client: %s", err.Error())
	}

	err = db.Ping()

	if err != nil {
		return nil, fmt.Errorf("unable to connect to `mysql`: %s", err.Error())
	}

	return &Client{db}, nil
}

// LogStats genera log con estadísticas de la conexión
func (c *Client) LogStats(logger *zap.Logger) {
	stats := c.Stats()

	logger.Info("mysql stats",
		zap.Int("idle", stats.Idle),
		zap.Int("inUse", stats.InUse),
		zap.Int("maxOpenConnections", stats.MaxOpenConnections),
		zap.Int("openConnections", stats.OpenConnections),
		zap.Int64("maxIdleClosed", stats.MaxIdleClosed),
		zap.Int64("maxLifetimeClosed", stats.MaxLifetimeClosed),
		zap.Int64("waitCount", stats.WaitCount),
		zap.Duration("waitCount", stats.WaitDuration),
	)
}
