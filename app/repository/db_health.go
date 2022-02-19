package repository

type Health struct {
	Status       string
	QueryCount   int
	SlaveRunning bool
}

// HealthRepository ...
type HealthRepository interface {
	RdbHealthCheck() (h *Health, err error)
}
