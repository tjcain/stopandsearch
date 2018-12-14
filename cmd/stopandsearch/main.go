package main

import "log"

func main() {
	// get config
	conf, err := server.GetConfigFromFile()
	if err != nil {
		log.Fatal(err)
	}

	// start server
	srv := server.New(conf)
	if err = server.Configure(); err != nil {
		log.Fatal(err)
	}

}
