package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fernandocao/ideazmxgoweb/internal/config"
	"github.com/fernandocao/ideazmxgoweb/internal/handlers"
	"github.com/fernandocao/ideazmxgoweb/internal/render"

	"github.com/alexedwards/scs/v2"
)

const puerto = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when is in production
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

	fmt.Println("Iniciando servidor en el puerto", puerto)

	srv := &http.Server{
		Addr:    puerto,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
