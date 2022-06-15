package main

import (
	"github.com/gorilla/mux"

	"IGISBackEnd/apis"

	"log"
	"net/http"
	"time"
)

func init() {
	log.Println("IGIS IDP Platform Backend Starting at http://127.0.0.1")
}

func main() {
	r := mux.NewRouter()

	// api landing page
	r.HandleFunc("/", apis.ServeLanding).
		Methods("GET")

	// api v1 subrouter
	sV1 := r.PathPrefix("/api/v1").Subrouter()
	sV1.HandleFunc("/asset", apis.ServeAssetWhole).
		Methods("GET")
	sV1.HandleFunc("/debt", apis.ServeDebtWhole).
		Methods("GET")
	sV1.HandleFunc("/macro", apis.ServeMacroWhole).
		Methods("GET")
		// Queries("asset", "{assetName}")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
