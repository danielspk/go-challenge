package domain

//RoutingService interface de servicio de cálculo de rutas
type RoutingService interface {
	GetRoute(*Point, *Point) (float32, error)
}
