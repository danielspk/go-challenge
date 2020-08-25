package application

import "challenge.com/challenge/internal/office/domain"

// FindByIdService estructura de servicio de búsqueda de una sucursal
type FindByIdService struct {
	Repository domain.Repository
}

// NewFindByIdService crea un FindByIdService
func NewFindByIdService(repository domain.Repository) *FindByIdService {
	return &FindByIdService{
		Repository: repository,
	}
}

// Execute corre el servicio
func (s *FindByIdService) Execute(cmd *FindByIdCommand) (*domain.Office, error) {
	office, err := s.Repository.FindById(cmd.ID)

	if err != nil {
		return nil, err
	}

	return office, nil
}
