package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// mapping methods is possible with httprouter
	router.ServeFiles("/static/*filepath", http.Dir("/home/jbpratt/static"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
