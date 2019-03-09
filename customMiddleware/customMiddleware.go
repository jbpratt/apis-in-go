package main

import (
	"fmt"
	"net/http"
)

// middleware func that accepts a handler and returns a handler
// embedding the main handler logic in it
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// pass control back to the handler
		// ServeHTTP allows handler to execute the logic that is in mainLogic
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// business logic
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("OK"))
}

func main() {
	// HandlerFunc returns HTTP handler
	// create handler func by passing the main handler func to http.HandlerFunc()
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	http.ListenAndServe(":8000", nil)
}
