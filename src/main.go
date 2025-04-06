package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(0)
	log.Print("Started!")

	mux := http.NewServeMux()
	mapRoutes(mux)

	e := http.ListenAndServe(":8080", mux)
	if e != nil {
		log.Fatal(e)
	}
}
