package main

import (
	"net/http"

	"github.com/Spx1010/web/pkg/config"
	"github.com/Spx1010/web/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/About", handlers.Repo.About)

	return mux

}
