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
	// process parameters
	yearFrom, err := strconv.Atoi(params["yearFrom"])
	if err != nil {
		log.Println("QryDebt :: yearFrom not integer")
	}
	yearUntil, err := strconv.Atoi(params["yearUntil"])
	if err != nil {
		log.Println("QryDebt :: yearUntil not integer")
	}
	qry := ReqIDPDebt{
		YearFrom:  yearFrom,
		YearUntil: yearUntil,
	}

	packet, _ := json.Marshal(procDebtQry(qry, d))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func procDebtQry(rq ReqIDPDebt, d IDPDebt) []debts {
	// reads debt query from ReqIDPDebt struct
	// process query by scanning for conditions
	// O(n) * conditions
	var sendPacket = []debts{}
	for _, row := range d.Data {
		y, _ := strconv.Atoi(row.LoanDate[:4])
		if (y >= rq.YearFrom) && (y <= rq.YearUntil) {
			sendPacket = append(sendPacket, row)
		}
	}
	return sendPacket
}
