package config

import (
	"fmt"
	"log"
	"task-management-system/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	host := "localhost"
	user := "postgres"
	password := "postgres"
	dbName := "task-management-system"
	port := "5432"
	sslMode := "disable"
	timezone := "Asia/Jakarta"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		host, user, password, dbName, port, sslMode, timezone,
	)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Gagal koneksi ke database: %v", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)
		sqlDB.SetConnMaxLifetime(time.Hour * 1)
	}

	log.Println("Koneksi database PostgresSQL berhassil.")
	return db
}

func AutoMigrate(db *gorm.DB) error {
	err := db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (
				SELECT 1
				FROM pg_type
				WHERE typname = 'task_status'
			) THEN
				CREATE TYPE task_status AS ENUM (
					'PENDING',
					'IN_PROGRESS',
					'COMPLETED'
				);
			END IF;
		END $$;
	`).Error

	if err != nil {
		return err
	}

	return db.AutoMigrate(&models.UserModel{}, &models.TaskModel{})
}
