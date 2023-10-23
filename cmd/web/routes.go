package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"net/http"

	//pat router
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AllCache) http.Handler{
	//mux:=pat.New()

	//mux.Get("/",http.HandlerFunc(handlers.Repo.Home))

	mux:=chi.NewRouter()

	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/",handlers.Repo.Home)

	return mux
}



//chi for routing-> has built in middleware-> allows process requests as they come
//has user has necessary rights to see some pages?
//Middleware acts as a bridge or intermediary that processes and potentially modifies data or behavior as it flows through the application.