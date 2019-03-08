package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Processing request...")
	w.Write([]byte("OK"))
	log.Println("Finished processing request")
}

func main() {
	r := mux.NewRouter()
	// attach handler called mainLogic
	r.HandleFunc("/", mainLogic)
	// wrapping router in handlers.LoggingHandler middleware
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	// then pass handler to ListenAndServe
	http.ListenAndServe(":8000", loggedRouter)
}
