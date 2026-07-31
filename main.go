package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	AppVersion string    `json:"app_version"`
	Color      string    `json:"color"`
	Hostname   string    `json:"hostname"`
	Time       time.Time `json:"time"`
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Читаем версию и цвет из переменных окружения
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "v1.0"
	}
	color := os.Getenv("APP_COLOR")
	if color == "" {
		color = "green"
	}

	hostname, _ := os.Hostname()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			AppVersion: version,
			Color:      color,
			Hostname:   hostname,
			Time:       time.Now(),
		})
	})

	http.ListenAndServe(":8080", r)
}
