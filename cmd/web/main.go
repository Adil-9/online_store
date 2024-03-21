package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	applicationlayer "github.com/Adil-9/online_store/internal/ApplicaitonLayer"
	dblayer "github.com/Adil-9/online_store/internal/DBlayer"
)

func main() {
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Println("Logger initiated")

	wh := &applicationlayer.WebHandler{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	db, err := dblayer.ConnectDB() //connection to database
	if err != nil {
		wh.ErrorLog.Fatal(err)
	}
	wh.InfoLog.Println("Database initiated")
	defer db.Close()

	wh.MustLoadHadnler(db) //creating database, service layer and parsing templates

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  wh.HandleRoutes(),
	}

	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = srv.ListenAndServe()
		if err != nil {
			errorLog.Fatal(err)
		}
		defer srv.Close()
	}()

	for range c {
		err := dblayer.CloseDB(db)
		if err != nil {
			wh.ErrorLog.Println("error closing database/migrating down", err)
		}
		wh.InfoLog.Println("Closing database connection")
		return
	}

}
