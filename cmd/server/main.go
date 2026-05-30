package main

import (
	"net/http"
	"task-management-system/config"
	delivery "task-management-system/internal/delivery/http"
	"task-management-system/internal/models"
	"task-management-system/internal/repository"
	"task-management-system/internal/usecase"

	"github.com/gin-gonic/gin"
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
