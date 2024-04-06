package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"pos_system/internal/config"
)

func ConnectDB(cfg config.DBConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
