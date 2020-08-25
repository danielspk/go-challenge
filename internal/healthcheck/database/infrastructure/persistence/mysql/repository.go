package mysql

import "database/sql"

// Repository estructura de repositorio MySQL
type Repository struct {
	DB *sql.DB
}

// NewRepository crea un Repository
func NewRepository(database *sql.DB) *Repository {
	return &Repository{
		DB: database,
	}
}

// Ping a la base de datos
func (r *Repository) Ping() error {
	return r.DB.Ping()
}
