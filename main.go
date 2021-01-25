package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/condezero/go-crud-api/internal/database"
	"github.com/condezero/go-crud-api/internal/server"
)

func main() {
	serv, err := server.New("5000")
	if err != nil {
		log.Fatal(err)
	}

	d := database.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	go serv.Start()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	serv.Close()
	database.Close()

}
