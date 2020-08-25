package application

import "challenge.com/challenge/internal/office/domain"

// CreateService estructura de servicio de creación de nueva sucursal
type CreateService struct {
	Repository domain.Repository
}

// NewCreateService crea un CreateService
func NewCreateService(repository domain.Repository) *CreateService {
	return &CreateService{
		Repository: repository,
	}
}

// Execute corre el servicio
func (s *CreateService) Execute(cmd *CreateCommand) (*domain.Office, []string, error) {
	office := &domain.Office{
		Address:   cmd.Address,
		Latitude:  cmd.Latitude,
		Longitude: cmd.Longitude,
	}

	messages, err := ValidateCreate(office)

	if err != nil {
		return nil, messages, err
	}

	newID, err := s.Repository.Save(office)

	if err != nil {
		return nil, nil, err
	}

	office.ID = newID

	return office, nil, nil
}
