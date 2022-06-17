package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ServeDebtWhole(d IDPDebt, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// process year slider
	yearFrom, err := strconv.Atoi(params["yearFrom"])
	if err != nil {
		log.Println("QryDebt :: yearFrom not integer")
	}
	yearUntil, err := strconv.Atoi(params["yearUntil"])
	if err != nil {
		log.Println("QryDebt :: yearUntil not integer")
	}

	var sendPacket = []debts{}
	for _, row := range d.Data {
		y, _ := strconv.Atoi(row.LoanDate[:4])
		if (y >= yearFrom) && (y <= yearUntil) {
			sendPacket = append(sendPacket, row)
		}
	}
	packet, _ := json.Marshal(sendPacket)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}
