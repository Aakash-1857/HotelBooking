package main

import (
	"net/http"

	"github.com/Aakash-1857/HotelBooking/pkg/config"
	"github.com/Aakash-1857/HotelBooking/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	mux:=chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	return mux
}