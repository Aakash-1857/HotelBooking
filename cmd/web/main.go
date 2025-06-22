package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Aakash-1857/HotelBooking/pkg/config"
	"github.com/Aakash-1857/HotelBooking/pkg/handlers"
	"github.com/Aakash-1857/HotelBooking/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
// main is the main function
func main() {
	app.InProduction = false
	var app config.AppConfig
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache")
	}
	app.TemplateCache = tc

	render.NewTemplate(&app)
	repo:=handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	//_ = http.ListenAndServe(portNumber, nil)
	serve := &http.Server{
		Addr : portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	if err!=nil{
		log.Fatal(err)
	}
}

