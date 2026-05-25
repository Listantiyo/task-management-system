package main

import (
	"encoding/json"
	"net/http"
	"task-management-system/config"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	_ "task-management-system/docs"

	"github.com/go-chi/chi/v5"
)

// @title 		Task Management System API
// @version 	1.0
// @description Ini adalah API untuk Task Management System.
// @host		localhost:8080
// @BasePath	/
func main() {
	config.InitDB();

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Get("/ping", PingHandler)


	srv := &http.Server{
		Addr: ":8080",
		Handler: r,
	}

	srv.ListenAndServe();
}

// PingHandler
// @Summary      ping test
// @Description  Menguji dengan menjalankan ping ke API
// @Tags         ping
// @Produce      json
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func PingHandler(w http.ResponseWriter, req *http.Request) {
		resp := map[string]string{
			"message": "pong",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}