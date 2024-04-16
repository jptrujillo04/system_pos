package location

import (
	"errors"
	"gorm.io/gorm"
	"log"
)

type RepositoryLocation interface {
	GetAllCountries(output interface{}) error
	SaveCountry(country *Country) error
	UpdateCountry(country *Country) error
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
	log.Print("initial transaction create country")
	tx := r.DBRepository.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := r.DBRepository.Create(country).Error; err != nil {
		log.Print("error en el repository create country. error: ", err)
		tx.Rollback()
		return errors.New("error repository create the country")
	}
	defer func() {
		if r := recover(); r != nil {
			log.Print("Recovered in panic of repository create country", r)
			tx.Rollback()
			panic(r)
		}
	}()
	log.Print("finish transaction create country")
	return tx.Commit().Error
}

func (r Repository) UpdateCountry(country *Country) error {
	log.Print("initial transaction update Country")
	tx := r.DBRepository.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	if err := r.DBRepository.Save(country).Error; err != nil {
		log.Print("error in repository update Country error: ", err)
		tx.Rollback()
		return errors.New("error repository update the country")
	}
	defer func() {
		if r := recover(); r != nil {
			log.Print("recovered in panic of repository update country", r)
			tx.Rollback()
			panic(r)
		}
	}()
	log.Print("finish transaction update country")
	return tx.Commit().Error
}
