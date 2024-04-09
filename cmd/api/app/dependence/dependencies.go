package dependence

import (
	"gorm.io/gorm"
	"log"
	"pos_system/internal/config"
	"pos_system/internal/database"
	"pos_system/internal/location"
)

type Dependencies struct {
	LocationRepository location.Repository
}

func NewDependencies() (*Dependencies, error) {
	db, err := ConnectionDataBase()
	if err != nil {
		return nil, err
	}
	dbRepository := location.NewRepositoryLocation(db)
	return &Dependencies{
		LocationRepository: dbRepository,
	}, nil
}

func ConnectionDataBase() (*gorm.DB, error) {
	dbConfig := config.ReadDBConfig()
	db, err := database.ConnectDB(dbConfig)
	if err != nil {
		log.Println("Error conectando a la base de datos:", err)
	}
	return db, nil
}
