package mysql

import (
	"challenge.com/challenge/internal/office/domain"
	"database/sql"
	"errors"
)

const (
	queryFindAll  = `SELECT id, address, latitude, longitude FROM offices`
	queryFindById = `SELECT id, address, latitude, longitude FROM offices WHERE id = ?`
	queryInsert   = `INSERT INTO offices (address, latitude, longitude) VALUES (?, ?, ?)`
)

//Repository estructura de repositorio MySQL
type Repository struct {
	DB *sql.DB
}

// NewRepository crea un Repository
func NewRepository(database *sql.DB) *Repository {
	return &Repository{
		DB: database,
	}
}

// FindAll busca todas las sucursales
func (r *Repository) FindAll() ([]*domain.Office, error) {
	var offices []*domain.Office

	results, err := r.DB.Query(queryFindAll)

	if err != nil {
		return nil, err
	}

	for results.Next() {
		var office domain.Office

		err := results.Scan(
			&office.ID, &office.Address,
			&office.Latitude, &office.Longitude,
		)

		if err != nil {
			return nil, err
		}

		offices = append(offices, &office)
	}

	return offices, nil
}

// FindById busca una sucursal por ID
func (r *Repository) FindById(id uint64) (*domain.Office, error) {
	var office domain.Office

	err := r.DB.QueryRow(queryFindById, id).Scan(
		&office.ID, &office.Address,
		&office.Latitude, &office.Longitude,
	)

	if err != nil {
		return nil, err
	}

	return &office, nil
}

// Save persiste una nueva sucursal
func (r *Repository) Save(office *domain.Office) (uint64, error) {
	if office.ID != 0 {
		return 0, errors.New("office update not implemented")
	}

	stmt, err := r.DB.Prepare(queryInsert)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(office.Address, office.Latitude, office.Longitude)

	if err != nil {
		return 0, err
	}

	newID, _ := result.LastInsertId()

	return uint64(newID), nil
}
