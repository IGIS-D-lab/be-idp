package main

import (
	"IGISBackEnd/apis"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	log.Println("IGIS IDP Platform Backend Starting at http://127.0.0.1:8080")
	log.Println("RowCountTest: ", apis.TEST_URL_ROW)
	log.Println("AssetTest: ", apis.TEST_URL_ASSET)
	log.Println("DebtTest: ", apis.TEST_URL_DEBT)
	log.Println("MacroTest: ", apis.TEST_URL_MACRO)
}

func main() {
	r := mux.NewRouter()

	// Pre-load data
	d := apis.MntData()

	// api landing page
	r.HandleFunc("/", apis.ServeLanding).
		Methods("GET")

	// api v1 subrouter
	sV1 := r.PathPrefix("/api/v1").Subrouter()
	sV1.Path("/debt").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, w, r)
		}).
		Name("ServeDebt")

	sV1.HandleFunc("/asset", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeAssetWhole(d.Asset, w, r)
	}).
		Methods("GET").
		Queries(
			"strat", "{strat}",
		)

	sV1.HandleFunc("/macro", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeMacroWhole(d.Macro, w, r)
	}).
		Methods("GET").
		Queries(
			"commodity", "{commodity}",
			"yearFrom", "{yearFrom:[0-9]+}",
			"yearUntil", "{yearUntil:[0-9]+}",
		)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
