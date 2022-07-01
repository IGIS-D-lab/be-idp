package main

import (
	"IGISBackEnd/apis"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	initAscii := `
	___ ___ ___ ___     _   ___ ___ 
	|_ _/ __|_ _/ __|   /_\ | _ \_ _|
	 | | (_ || |\__ \  / _ \|  _/| | 
	|___\___|___|___/ /_/ \_\_| |___|
									 `
	fmt.Println(initAscii)
	log.Println("IGIS IDP Platform Backend Starting at http://127.0.0.1:8080")
}

func main() {
	// _, _ = logs.LogInit()
	r := mux.NewRouter()

	// Pre-load data
	d := apis.MntData()

	// api landing page
	r.HandleFunc("/", apis.ServeLanding).
		Methods("GET")

	// api v1 subrouter -  debt
	sV1Debt := r.PathPrefix("/api/v1/debt").Subrouter()
	sV1Debt.Path("/dataTable").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, 0, w, r)
		}).
		Methods("GET").
		Name("debt.datatable")
	sV1Debt.Path("/graphLeft").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, 1, w, r)
		}).
		Methods("GET").
		Name("debt.graphLeft")
	sV1Debt.Path("/graphRight").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, 2, w, r)
		}).
		Methods("GET").
		Name("debt.graphRight")

	// api v1 subrouter - model
	sV1Model := r.PathPrefix("/api/v1/model").Subrouter()
	sV1Model.Path("/info").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeModelInfo(d.ModelInfo, w, r)
		}).
		Methods("GET").
		Name("model information")
	sV1Model.Path("/coef").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeModelCoef(d.ModelCoef, w, r)
		}).
		Methods("GET").
		Name("model coefficients")
	sV1Model.Path("/pred").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeModelCalc(d.ModelCoef, d.Macro, w, r)
		}).
		Methods("GET").
		Name("model prediction")

	// api v1 subrouter all
	sV1 := r.PathPrefix("/api/v1").Subrouter()
	sV1.HandleFunc("/single", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeSingle(d.Debt, w, r)
	}).
		Methods("GET").
		Name("singleInfo")

	// TODO: EDIT Below
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

	// serve
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
