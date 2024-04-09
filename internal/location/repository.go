package location

import (
	"gorm.io/gorm"
)

type RepositoryLocation interface {
	GetAllContries(output interface{}) error
}

type Repository struct {
	DBRepository *gorm.DB
}

func NewRepositoryLocation(db *gorm.DB) Repository {
	return Repository{
		DBRepository: db,
	}
}

func (r Repository) GetAllContries(output interface{}) error {
	if err := r.DBRepository.Find(output).Error; err != nil {
		return err
	}
	return nil
}
