package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
	"grig/internal/handler/bus"
	"grig/internal/handler/group"
	"grig/internal/handler/journal"
	"grig/internal/handler/student"
	"grig/internal/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer)
	r.Use(middleware.Timeout(time.Second * 60))
	cors := cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:             300,
	})
	r.Use(cors.Handler)
	repo, err := repository.New()
	if err != nil {
		log.Fatalf("failed to initialize cause %+v", err)
	}
	studentHandler := student.New(repo)
	groupHandler := group.New(repo)
	journalHandler := journal.New(repo)
	busHandler := bus.New()
	r.Route("/api", func(r chi.Router) {
		r.Post("/bus/students/{id}", busHandler.Post)

		r.Get("/students/{id}", studentHandler.Get)
		r.Get("/students", studentHandler.GetAll)
		r.Put("/students", studentHandler.Put)
		r.Delete("/students/{id}", studentHandler.Delete)
		r.Post("/students", studentHandler.Post)

		r.Get("/groups/{id}", groupHandler.Get)
		r.Get("/groups", groupHandler.GetAll)
		r.Put("/groups", groupHandler.Put)
		r.Delete("/groups/{id}", groupHandler.Delete)
		r.Post("/groups", groupHandler.Post)

		r.Get("/journals/{id}", journalHandler.Get)
		r.Get("/journals", journalHandler.GetAll)
		r.Put("/journals", journalHandler.Put)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
