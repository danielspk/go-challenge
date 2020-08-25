package application

// CreateCommand estructura de tipo DTO
type CreateCommand struct {
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
