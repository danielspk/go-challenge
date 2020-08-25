package graphhopper

// RouteResponse estructura de tipo DTO
type RouteResponse struct {
	Paths []paths `json:"paths"`
}

// paths estructura de tipo DTO
type paths struct {
	Distance float32 `json:"distance"`
	Weight   float32 `json:"weight"`
	Time     uint64  `json:"time"`
}
