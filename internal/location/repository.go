package location

import (
	"database/sql"
)

type RepositoryLocation interface {
	GetAllContries(output interface{}) error
}

type Repository struct {
	DBRepository *sql.DB
}

func NewRepositoryLocation(db *sql.DB) Repository {
	return Repository{
		DBRepository: db,
	}
}

func (r Repository) GetAllContries(output interface{}) error {
	return nil
}
