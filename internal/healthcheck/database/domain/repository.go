package domain

//Repository interface de repositorio de servicios
type Repository interface {
	Ping() error
}
