package graphhopper

import (
	"challenge.com/challenge/internal/search/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Service estructura de servicio a Graph Hopper
type Service struct {
	ApiKey string
}

// NewService crea un Service
func NewService(apiKey string) *Service {
	return &Service{
		ApiKey: apiKey,
	}
}

// GetRoute consulta la distancia de una ruta entre dos puntos
func (s *Service) GetRoute(origin *domain.Point, destination *domain.Point) (float32, error) {
	url := fmt.Sprintf(
		"https://graphhopper.com/api/1/route?point=%v,%v&point=%v,%v&vehicle=car&locale=es&calc_points=false&key=%s",
		origin.Latitude, origin.Longitude, destination.Latitude, destination.Longitude, s.ApiKey,
	)

	log.Printf("graphhopper call to URL: %s", url)
	resp, err := http.Get(url)

	if err != nil {
		return 0, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Printf("graphhopper response: %s", bodyBytes)

	var routeResponse *RouteResponse

	err = json.Unmarshal(bodyBytes, &routeResponse)

	if err != nil {
		return 0, err
	}

	if len(routeResponse.Paths) < 1 {
		return 0, errors.New("empty paths")
	}

	distance := routeResponse.Paths[0].Distance

	if distance == 0 {
		return 0, errors.New("empty distance")
	}

	return distance, nil
}
