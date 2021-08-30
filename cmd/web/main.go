package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/urschmidt/bookings/pkg/config"
	"github.com/urschmidt/bookings/pkg/handlers"
	"github.com/urschmidt/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to false when in production
	appConfig.DevelopmentMode = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = !appConfig.DevelopmentMode

	appConfig.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	appConfig.TemplateCache = tc

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfig)

	fmt.Println(fmt.Sprintf("Starting Application on Port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&appConfig),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
