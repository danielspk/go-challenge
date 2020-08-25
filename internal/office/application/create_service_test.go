package application

import (
	"challenge.com/challenge/internal/office/infrastructure/persistence/memory"
	"testing"
)

func Test_correctOnCreate(t *testing.T) {
	repository := memory.NewRepository()
	service := NewCreateService(repository)
	command := &CreateCommand{Address: "testing address", Latitude: -34.123456, Longitude: -58.123456}

	office, details, err := service.Execute(command)

	if err != nil {
		t.Fail()
		return
	}

	if details != nil {
		t.Fail()
		return
	}

	if office.ID != 4 {
		t.Fail()
	}
}

func Test_failByEmptyAddressOnCreate(t *testing.T) {
	repository := memory.NewRepository()
	service := NewCreateService(repository)
	command := &CreateCommand{Address: "", Latitude: -34.123456, Longitude: -58.123456}

	_, details, err := service.Execute(command)

	if err == nil {
		t.Fail()
		return
	}

	if len(details) == 0 {
		t.Fail()
		return
	}

	if details[0] != "the address cannot be empty" {
		t.Fail()
	}
}
