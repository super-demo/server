package app

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgresClient *gorm.DB
)

func PostgresConnect(config *PostgresConfigs) (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", config.Host, config.User, config.Password, config.Database, config.Port, config.SSLMode)
	gormConfig := gorm.Config{
		SkipDefaultTransaction: true,
	}

	db, err := gorm.Open(postgres.Open(dsn), &gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Second * 10)

	err = sqlDB.Ping()
	if err != nil {
		return nil, err
	}

	SLog.Info("Connected to Postgresql")

	PostgresClient = db

	return PostgresClient, nil
}
