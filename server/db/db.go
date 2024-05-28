package db

import (
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"hospital-system/server/config"
)

var DB *gorm.DB

func OpenConnection(config config.Config) (*gorm.DB, error) {
	var err error
	connectionString := getDatabaseConnectionString(config.Database)

	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to connect to the Database, error %s", err.Error()))
		return nil, err
	}

	slog.Info("🚀 Connected Successfully to the Database")

	return DB, nil
}

func CloseConnection() error {
	defer func() {
		v := recover()
		if v != nil {
			panic(v)
		} else {
			slog.Info("Databased connection closed successfully")
		}
	}()

	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func getDatabaseConnectionString(config config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Host,
		config.Port,
		config.User,
		config.DB,
		config.Pwd,
	)
}
