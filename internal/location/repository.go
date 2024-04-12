package location

import (
	"errors"
	"gorm.io/gorm"
)

type RepositoryLocation interface {
	GetAllCountries(output interface{}) error
	SaveCountry(country *Country) error
}

type Repository struct {
	DBRepository *gorm.DB
}

func NewRepositoryLocation(db *gorm.DB) Repository {
	return Repository{
		DBRepository: db,
	}
}

func (r Repository) GetAllCountries(output interface{}) error {
	if err := r.DBRepository.Find(output).Error; err != nil {
		return errors.New("error repository find the country")
	}
	return nil
}

func (r Repository) SaveCountry(country *Country) error {
	tx := r.DBRepository.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := r.DBRepository.Create(country).Error; err != nil {
		errors.New("error repository save the country")
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()
	return tx.Commit().Error
}
