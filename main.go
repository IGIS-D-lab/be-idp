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
	log.Println("RowCountTest: http://localhost:8080/api/v1/debtRowCount?yearFrom=2000&yearUntil=2021&aumFrom=1&aumUntil=100000000000&debtFrom=1&debtUntil=1000000000000")
	log.Println("AssetTest: http://localhost:8080/api/v1/asset?strat=Core")
	log.Println("DebtTest: http://localhost:8080/api/v1/debt?yearFrom=2010&yearUntil=2020")
	log.Println("MacroTest: http://localhost:8080/api/v1/macro?commodity=kr1y&yearFrom=2010&yearUntil=2020")
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
	sV1.HandleFunc("/debtRowCount", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeDebtRowCount(d.Debt, w, r)
	}).
		Methods("GET").
		Queries(
			"yearFrom", "{yearFrom:[0-9]+}",
			"yearUntil", "{yearUntil:[0-9]+}",
			"aumFrom", "{aumFrom}",
			"aumUntil", "{aumUntil}",
			"debtFrom", "{debtFrom}",
			"debtUntil", "{debtUntil}",
		)
	sV1.HandleFunc("/asset", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeAssetWhole(d.Asset, w, r)
	}).
		Methods("GET").
		Queries(
			"strat", "{strat}",
		)
	sV1.HandleFunc("/debt", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeDebtWhole(d.Debt, w, r)
	}).
		Methods("GET").
		Queries(
			"yearFrom", "{yearFrom:[0-9]+}",
			"yearUntil", "{yearUntil:[0-9]+}",
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
