package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/guicattani/go-course/pkg/config"
	"github.com/guicattani/go-course/pkg/handlers"
	"github.com/guicattani/go-course/pkg/render"
)

var ac config.AppConfig
var session *scs.SessionManager

func main() {

	ac.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = ac.InProduction

	ac.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	ac.TemplateCache = tc

	repo := handlers.NewRepo(&ac)
	handlers.NewHandlers(repo)

	render.NewTemplates(&ac)

	http.HandleFunc("/", handlers.Repo.Home)

	fmt.Println("Listening on port 8080")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&ac),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
