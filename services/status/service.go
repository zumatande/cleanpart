package status

// Service interface defines Status service.
type Service interface {
	Status() error
}
