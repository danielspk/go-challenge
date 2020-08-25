package mysql

import (
	"github.com/golang-migrate/migrate/v4"
	mysqlDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// Migrate ejecuta las migraciones en la base de datos
func Migrate(mysqlClient *Client, sourceURL string, databaseName string, version uint) error {
	driver, err := mysqlDriver.WithInstance(mysqlClient.DB, &mysqlDriver.Config{})

	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(sourceURL, databaseName, driver)

	if err != nil {
		return err
	}

	err = m.Migrate(version)

	if err != nil && err.Error() != "no change" {
		return err
	}

	v, _, _ := m.Version()
	log.Printf("synchronized migration to version %d", v)

	return nil
}
