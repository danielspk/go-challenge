package domain

//Repository interface de repositorio de servicios
type Repository interface {
	FindAll() ([]*Office, error)
	FindById(uint64) (*Office, error)
	Save(*Office) (uint64, error)
}
