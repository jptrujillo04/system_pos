package dependence

import (
	"database/sql"
	"pos_system/internal/location"
)

type Dependencies struct {
	LocationRepository location.Repository
}

func NewDependencies(db *sql.DB) (*Dependencies, error) {
	dbRepository := location.NewRepositoryLocation(db)
	return &Dependencies{
		LocationRepository: dbRepository,
	}, nil
}
