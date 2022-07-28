package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/guicattani/handy-folk-finder/internal/config"
	"github.com/guicattani/handy-folk-finder/internal/driver"
	"github.com/guicattani/handy-folk-finder/internal/handlers"
	"github.com/guicattani/handy-folk-finder/internal/helpers"
	"github.com/joho/godotenv"
)

var ac config.AppConfig

var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Printf("Staring application on port :%s \n", os.Getenv("PORT_NUMBER"))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT_NUMBER")),
		Handler: routes(&ac),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {
	// change this to true when in production
	ac.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ac.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	ac.ErrorLog = errorLog

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(helpers.DbConfig())
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	repo := handlers.NewRepo(&ac, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&ac)

	return db, nil
}
