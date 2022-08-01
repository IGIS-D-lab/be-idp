package main

import (
	"IGISBackEnd/apis"
	"IGISBackEnd/orm"
	v2 "IGISBackEnd/v2"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
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

func routeLanding(rt *mux.Router, d apis.IDPDataSet) {
	// single fund request sub-endpoint
	rt.HandleFunc("/single", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeSingle(d.Debt, w, r)
	}).
		Methods("GET").
		Name("singleInfo")
	// asset(checklist sheet) sub-endpoint
	rt.HandleFunc("/asset", func(w http.ResponseWriter, r *http.Request) {
		apis.ServeAssetWhole(d.Asset, w, r)
	}).
		Methods("GET").
		Queries(
			"strat", "{strat}",
		)
}

func routeDebt(rt *mux.Router, d apis.IDPDataSet) {
	// dataTable sub-endpoint
	rt.Path("/dataTable").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, 0, w, r)
		}).
		Methods("GET").
		Name("debt.datatable")
	// graphLeft sub-endpoint
	rt.Path("/graphLeft").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, 1, w, r)
		}).
		Methods("GET").
		Name("debt.graphLeft")
	// graphRight sub-endpoint
	rt.Path("/graphRight").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeDebt(d.Debt, 2, w, r)
		}).
		Methods("GET").
		Name("debt.graphRight")
}

func routeModel(rt *mux.Router, d apis.IDPDataSet) {
	// model info sub-endpoint
	rt.Path("/info").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeModelInfo(d.ModelInfo, w, r)
		}).
		Methods("GET").
		Name("model information")
	// model coef sub-endpoint
	rt.Path("/coef").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeModelCoef(d.ModelCoef, w, r)
		}).
		Methods("GET").
		Name("model coefficients")
	// model prediction sub-endpoint
	rt.Path("/pred").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeModelCalc(d.ModelCoef, d.ModelInfo, d.Macro, w, r)
		}).
		Methods("GET").
		Name("model prediction")
}

func routeMacro(rt *mux.Router, db *redis.Client, d apis.IDPDataSet, du *apis.IDPDataSet) {
	rt.Path("/dataTable").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.ServeMacro(db, w, r)
		}).
		Methods("GET").Name("macro data")
	rt.Path("/update").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apis.UpdateMacro(&du.Macro, db, w, r)
		}).
		Methods("POST").
		Name("macro data update")
}

func routeMacro2(rt *mux.Router, rdb *redis.Client) {
	// data table
	rt.Path("/dataTable").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v2.GetMacro(rdb, w, r)
		}).
		Methods("GET").
		Name("macro data")
	rt.Path("/domUpdate").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v2.PostMacro(rdb, true, w, r)
		}).
		Methods("POST").
		Name("macro data update")
	rt.Path("/forUpdate").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v2.PostMacro(rdb, false, w, r)
		})
}

func routeMsg(rt *mux.Router, rdb *redis.Client) {
	rt.Path("/message").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v2.GetMessage(rdb, w, r)
		}).
		Methods("GET").
		Name("get message")
	rt.Path("/newMessage").
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			v2.PostMessage(rdb, w, r)
		}).
		Methods("POST").
		Name("post new message")
}

func main() {
	// _, _ = logs.LogInit()
	r := mux.NewRouter()
	database, err := orm.Conn("./token.json")
	if err != nil {
		log.Panicln(err)
	}

	// Pre-load data
	d := apis.MntData()

	// api landing page
	r.HandleFunc("/", apis.ServeLanding).Methods("GET")

	// api v1 subrouter all
	sV1 := r.PathPrefix("/api/v1").Subrouter()
	routeLanding(sV1, d)

	// api v1 subrouter -  debt
	sV1Debt := r.PathPrefix("/api/v1/debt").Subrouter()
	routeDebt(sV1Debt, d)

	// api v1 subrouter - model
	sV1Model := r.PathPrefix("/api/v1/model").Subrouter()
	routeModel(sV1Model, d)

	// api v1 subrouter - macro
	sV1Macro := r.PathPrefix("/api/v1/macro").Subrouter()
	routeMacro(sV1Macro, database, d, &d)

	sV2Macro := r.PathPrefix("/api/v2/macro").Subrouter()
	routeMacro2(sV2Macro, database)

	// api v2 subrouter - message
	sV2Message := r.PathPrefix("/api/v2/message").Subrouter()
	routeMsg(sV2Message, database)

	// serve
	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
