package domain

// Office entidad de dominio que modela una sucursal
type Office struct {
	ID        uint64  `json:"id"`
	Address   string  `json:"address"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}
