package main

import (
	"challenge.com/challenge/internal/search/infrastructure/graphhopper"
	"challenge.com/challenge/internal/shared/infrastructure"
	"challenge.com/challenge/internal/shared/infrastructure/http/api"
	"challenge.com/challenge/internal/shared/infrastructure/logger"
	persistence "challenge.com/challenge/internal/shared/infrastructure/persistence/mysql"
	"fmt"
	"go.uber.org/zap"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	migrationUpVersion = 2  // migration target version
	delayDatabaseStats = 30 // delay of database stats in seconds
)

func main() {
	validateEnv()

	container := newContainer()

	defer func(container *infrastructure.Container) {
		_ = container.Database.Close()
	}(container)

	runMigrations(container.Database)
	runStatsDatabase(container.Database, container.Logger)
	runServer(container)
}

// validateEnv valida la presencia de las variables de entorno requeridas
func validateEnv() {
	envs := []string{"API_PORT", "DATABASE_USER", "DATABASE_PASSWORD", "DATABASE_HOST_PORT", "DATABASE_NAME", "GRAPHHOPPER_APY_KEY"}

	for _, envKey := range envs {
		_, exists := os.LookupEnv(envKey)

		if exists == false {
			log.Fatalf("missing environment variable `%s`", envKey)
		}
	}
}

// newContainer crea un contenedor de servicios
func newContainer() *infrastructure.Container {
	zapLogger, err := logger.NewZapInstance()

	if err != nil {
		log.Fatalf("Main Error: %s", err.Error())
	}

	database, err := persistence.NewMysqlClient(
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s",
			os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST_PORT"), os.Getenv("DATABASE_NAME"),
		),
	)

	if err != nil {
		log.Fatalf("Main Error: %s", err.Error())
	}

	graphHopper := graphhopper.NewService(os.Getenv("GRAPHHOPPER_APY_KEY"))

	return &infrastructure.Container{
		Logger:      zapLogger,
		Database:    database,
		GraphHopper: graphHopper,
	}
}

// runMigrations ejecuta las migraciones de base de datos
func runMigrations(database *persistence.Client) {
	err := persistence.Migrate(database, "file://scripts/migrations", os.Getenv("DATABASE_NAME"), migrationUpVersion)

	if err != nil {
		log.Fatalf("Main Error: %s", err.Error())
	}
}

// runStatsDatabase ejecuta una goroutine para registrar periódicamente estadísticas de la base de datos
func runStatsDatabase(database *persistence.Client, logger *zap.Logger) {
	go func(database *persistence.Client) {
		for {
			database.LogStats(logger)

			time.Sleep(delayDatabaseStats * time.Second)
		}
	}(database)
}

// runServer ejecuta el web server de la API
func runServer(container *infrastructure.Container) {
	port, err := strconv.ParseUint(os.Getenv("API_PORT"), 10, 16)

	if err != nil {
		log.Fatalf("Main Error: %s", err.Error())
	}

	server := api.Factory(uint16(port), container)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Main Error: %s", err.Error())
	}
}
