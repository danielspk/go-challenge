package memory

import (
	"challenge.com/challenge/internal/office/domain"
	"errors"
)

//Repository estructura de repositorio en memoria
type Repository struct {
	Rows []*domain.Office
}

// NewRepository crea un Repository con datos de prueba
func NewRepository() *Repository {
	return &Repository{
		Rows: []*domain.Office{
			{ID: 1, Address: "test address 1", Latitude: -34.545773, Longitude: -58.471413},
			{ID: 2, Address: "test address 2", Latitude: -34.611839, Longitude: -58.454481},
			{ID: 3, Address: "test address 3", Latitude: -34.669758, Longitude: -58.455486},
		},
	}
}

// FindAll busca en memoria todas las sucursales
func (r *Repository) FindAll() ([]*domain.Office, error) {
	return r.Rows, nil
}

// FindById busca en memoria una sucursal por ID
func (r *Repository) FindById(id uint64) (*domain.Office, error) {
	if uint64(len(r.Rows)) < id {
		return nil, errors.New("office not exists")
	}

	office := r.Rows[id-1]

	return office, nil
}

// Save persiste en memoria una nueva sucursal
func (r *Repository) Save(office *domain.Office) (uint64, error) {
	lastID := uint64(len(r.Rows)) + 1

	office.ID = lastID

	r.Rows = append(r.Rows, office)

	return lastID, nil
}
