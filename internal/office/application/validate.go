package application

import (
	"challenge.com/challenge/internal/office/domain"
	"errors"
)

// ValidateCreate valida la creación de una sucursal
func ValidateCreate(office *domain.Office) ([]string, error) {
	var messages []string

	if err := validateAddress(office.Address); err != nil {
		messages = append(messages, err.Error())
	}

	if err := validateLatitude(office.Latitude); err != nil {
		messages = append(messages, err.Error())
	}

	if err := validateLongitude(office.Longitude); err != nil {
		messages = append(messages, err.Error())
	}

	if len(messages) > 0 {
		return messages, errors.New("validation with errors")
	}

	return nil, nil
}

// validateAddress valida la dirección de una sucursal
func validateAddress(address string) error {
	addressLen := len(address)

	if addressLen == 0 {
		return errors.New("the address cannot be empty")
	} else if addressLen > 160 {
		return errors.New("the address must not exceed 160 characters")
	}

	return nil
}

// validateLatitude valida la latitud de una sucursal
func validateLatitude(latitude float32) error {
	if latitude < -90 || latitude > 90 {
		return errors.New("latitude must be between -90 and 90")
	}

	return nil
}

// validateLongitude valida la longitud de una sucursal
func validateLongitude(longitude float32) error {
	if longitude < -180 || longitude > 180 {
		return errors.New("longitude must be between -180 and 180")
	}

	return nil
}
