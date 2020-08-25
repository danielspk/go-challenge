package domain

import officeDomain "challenge.com/challenge/internal/office/domain"

// Point estructura de tipo DTO
type Point struct {
	Latitude  float32
	Longitude float32
}

// OfficeByProximity entidad virtual de respuesta
type OfficeByProximity struct {
	Distance             float32 `json:"distance"`
	*officeDomain.Office `json:"office"`
}
