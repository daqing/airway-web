package orm

import (
	"errors"
	"log"

	"github.com/daqing/airway/lib/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var __gormDB__ *gorm.DB

var ErrNotSetup = errors.New("database is not setup yet")

func Setup() error {
	pgURL, err := utils.GetEnv("AIRWAY_PG_URL")
	if err != nil {
		// skip setup database if no pg url provided
		return nil
	}

	__gormDB__, err = gorm.Open(postgres.Open(pgURL), &gorm.Config{})

	if err != nil {
		log.Printf("Failed to open database from gorm: %v", err)
		return err
	}

	return nil
}

func DB() *gorm.DB {
	pgURL, err := utils.GetEnv("AIRWAY_PG_URL")
	if err != nil && pgURL == "" {
		return nil
	}

	if __gormDB__ == nil {
		panic(ErrNotSetup)
	}

	return __gormDB__
}
