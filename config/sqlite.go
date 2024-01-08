package config

import (
	"os"

	"github.com/diogenesOliver/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQlite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbpath := "./db/main.db"
	_, err := os.Stat(dbpath)

	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbpath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("Sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
