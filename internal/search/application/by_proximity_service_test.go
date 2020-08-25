package application

import (
	"challenge.com/challenge/internal/office/infrastructure/persistence/memory"
	"challenge.com/challenge/internal/search/infrastructure/graphhopper"
	"os"
	"testing"
)

func Test_correctOnByProximity(t *testing.T) {
	repository := memory.NewRepository()
	graphhopperService := graphhopper.NewService(os.Getenv("GRAPHHOPPER_APY_KEY"))

	service := NewByProximityService(repository, graphhopperService)
	command := &ByProximityCommand{Latitude: -34.622678, Longitude: -58.478349}

	officeByProximity, err := service.Execute(command)

	if err != nil {
		t.Fail()
		return
	}

	if officeByProximity.Office.ID != 2 {
		t.Fail()
		return
	}

	distanceExpected := float32(4148.816)

	if officeByProximity.Distance < distanceExpected*0.75 || officeByProximity.Distance > distanceExpected*1.25 {
		t.Fail()
	}
}
