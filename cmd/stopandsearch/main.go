package main

import (
	"log"
	"net/http"

	"stopandsearch/pkg/server"
)

func main() {
	// get config
	conf := server.NewConfig()

	// start server
	srv := server.New(conf)
	if err := srv.Configure(); err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(srv.Config.Bind, srv.NewRouter())

}
