package apis

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IGIS Debt Platform landing page\n")
}

func ServeAssetWhole(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./asset/idpChecklist.json")
	if err != nil {
		log.Println("Checklist :: ", err)
	} else {
		log.Println("Checklist :: Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(byteVal)
}

func ServeDebtWhole(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./asset/idpDebt.json")
	if err != nil {
		log.Println("Debt :: ", err)
	} else {
		log.Println("Debt :: Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(byteVal)
}

func ServeMacroWhole(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./asset/idpMacro.json")
	if err != nil {
		log.Println("Macro :: ", err)
	} else {
		log.Println("Macro :: Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	// mount json
	w.Write(byteVal)
}
