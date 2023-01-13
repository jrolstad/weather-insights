package repositories

import (
	"github.com/jrolstad/weather-insights/internal/logging"
	"github.com/jrolstad/weather-insights/internal/models"
)

type ObservationRepository interface {
	Save(data []*models.Observation) error
}

func NewObservationRepository() ObservationRepository {
	return &ConsoleObservationRepository{}
}

type ConsoleObservationRepository struct {
}

func (r *ConsoleObservationRepository) Save(data []*models.Observation) error {
	logging.LogInfo("Saving Data", "data", data)
	return nil
}
