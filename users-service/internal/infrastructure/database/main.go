package database

import (
	"fmt"
	"users-service/internal/infrastructure/database/repository"
	"users-service/pkg"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(host, user, password, dbname, port string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=America/Sao_Paulo",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		pkg.Logger.Error("Failed to connect to database", zap.Error(err))
		return nil, err
	}

	err = db.AutoMigrate(&repository.UserModel{})
	if err != nil {
		pkg.Logger.Error("Failed to migrate user", zap.Error(err))
		return nil, err
	}

	return db, nil
}
