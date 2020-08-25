package application

import (
	officeDomain "challenge.com/challenge/internal/office/domain"
	"challenge.com/challenge/internal/search/domain"
	"errors"
	"sync"
)

// ByProximityService estructura de servicio de búsqueda de sucursal cercana
type ByProximityService struct {
	Repository officeDomain.Repository
	Service    domain.RoutingService
}

// NewByProximityService crea un ByProximityService
func NewByProximityService(
	repository officeDomain.Repository, service domain.RoutingService,
) *ByProximityService {
	return &ByProximityService{
		Repository: repository,
		Service:    service,
	}
}

// Execute corre el servicio
func (s *ByProximityService) Execute(cmd *ByProximityCommand) (*domain.OfficeByProximity, error) {
	offices, _ := s.Repository.FindAll()

	var bestOfficeByProximity *domain.OfficeByProximity
	var wg sync.WaitGroup
	ch := make(chan *domain.OfficeByProximity, len(offices))

	for _, office := range offices {
		wg.Add(1)

		go func(cmd *ByProximityCommand, office *officeDomain.Office, ch chan *domain.OfficeByProximity) {
			defer wg.Done()

			origin := &domain.Point{Latitude: cmd.Latitude, Longitude: cmd.Longitude}
			destination := &domain.Point{Latitude: office.Latitude, Longitude: office.Longitude}
			distance, err := s.Service.GetRoute(origin, destination)

			if err != nil {
				ch <- nil
				return
			}

			ch <- &domain.OfficeByProximity{
				Distance: distance,
				Office:   office,
			}
		}(cmd, office, ch)
	}

	wg.Wait()
	close(ch)

	for result := range ch {
		if result == nil {
			continue
		}

		if bestOfficeByProximity == nil || result.Distance < bestOfficeByProximity.Distance {
			bestOfficeByProximity = result
		}
	}

	if bestOfficeByProximity == nil {
		return nil, errors.New("could not find proximity offices")
	}

	return bestOfficeByProximity, nil
}
