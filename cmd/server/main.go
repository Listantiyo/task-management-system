package main

import (
	"net/http"
	"task-management-system/config"
	delivery "task-management-system/internal/delivery/http"
	customType "task-management-system/internal/delivery/types"
	"task-management-system/internal/models"
	"task-management-system/internal/repository"
	"task-management-system/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// @title 		Task Management System API
// @version 	1.0
// @description Ini adalah API untuk Task Management System.
// @host		localhost:8080
// @BasePath 	/api

// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
// @description                Masukkan token dengan format: Bearer <token_kamu>
func main() {
	db := config.InitDB()
	db.AutoMigrate(&models.UserModel{})

	r := gin.Default()

	// Register custom validator
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterValidation("enum", func(fl validator.FieldLevel) bool {
		status, ok := fl.Field().Interface().(customType.TaskStatus)
		if !ok {
			return false
		}
		return status.IsValid()
	})

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)

	handlers := delivery.Handlers{
		User: delivery.NewUserHandler(userUsecase),
	}

	r = delivery.NewRouter(handlers)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	srv.ListenAndServe()
}
