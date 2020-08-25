package application

import "challenge.com/challenge/internal/healthcheck/database/domain"

// PingService estructura de servicio de health check de base de datos
type PingService struct {
	Repository domain.Repository
}

// NewPingService crea un PingService
func NewPingService(repository domain.Repository) *PingService {
	return &PingService{
		Repository: repository,
	}
}

// Execute corre el servicio
func (s *PingService) Execute() error {
	return s.Repository.Ping()
}
