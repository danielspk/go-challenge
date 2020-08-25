package graphhopper

import (
	"challenge.com/challenge/internal/search/domain"
	"os"
	"testing"
)

func Test_correctGetRoute(t *testing.T) {
	service := NewService(os.Getenv("GRAPHHOPPER_APY_KEY"))

	origin := &domain.Point{
		Latitude:  -34.603832,
		Longitude: -58.381607,
	}

	destination := &domain.Point{
		Latitude:  -34.703832,
		Longitude: -58.481607,
	}

	distance, err := service.GetRoute(origin, destination)

	if err != nil {
		t.Fail()
		return
	}

	if distance <= 0 {
		t.Fail()
	}
}
