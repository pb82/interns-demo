package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	var name = "World"
	if value, found := os.LookupEnv("HELLO_WHO"); found {
		name = value
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println(fmt.Sprintf("request received from %s", r.RemoteAddr))
	fmt.Fprintf(w, fmt.Sprintf("Hello %s", name))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)

	port := 3000
	fmt.Println(fmt.Sprintf("starting server on port %d", port))

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
