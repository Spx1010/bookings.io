package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Spx1010/web/pkg/config"
	"github.com/Spx1010/web/pkg/handlers"
	"github.com/Spx1010/web/pkg/render"
	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

func main() {

	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	fmt.Println("Server running ...")
	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/About", handlers.Repo.About)

	//http.ListenAndServe(portNumber, nil)
	serve := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)

}
