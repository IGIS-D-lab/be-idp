package apis

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func ServeDebtRowCount(d IDPDebt, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// process year slider
	yearFrom, err := strconv.Atoi(params["yearFrom"])
	if err != nil {
		log.Println("yearFrom not integer")
	}
	yearUntil, err := strconv.Atoi(params["yearUntil"])
	if err != nil {
		log.Println("yearUntil not integer")
	}
	err = verifyRowCountYear(yearFrom, yearUntil)

	// process AUM slider
	aumFrom, err := strconv.Atoi(params["aumFrom"])
	if err != nil {
		log.Println("aumFrom not integer", params["aumFrom"])
	}
	aumUntil, err := strconv.Atoi(params["aumUntil"])
	if err != nil {
		log.Println("aumUntil not integer", params["aumUntil"])
	}

	// process Debt slider
	debtFrom, err := strconv.Atoi(params["debtFrom"])
	if err != nil {
		log.Println("debtFrom not integer")
	}
	debtUntil, err := strconv.Atoi(params["debtUntil"])
	if err != nil {
		log.Println("debtUntil not integer")
	}

	sendPacket := IDPRowCount{
		RowCount: debtRowCount(d, yearFrom, yearUntil, aumFrom, aumUntil, debtFrom, debtUntil),
	}
	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	// mount json
	w.Write(packet)
}

func verifyRowCountYear(yearFrom, yearUntil int) error {
	if yearFrom <= yearUntil {
		return nil
	} else {
		return errors.New("yearFrom is bigger than yearUntil")
	}
}

func isWithIn(s, e int, val string) bool {
	v := strings.Replace(val, ",", "", -1)
	res, _ := strconv.Atoi(v)
	if (res >= s) && (res <= e) {
		return true
	} else {
		return false
	}
}

func debtRowCount(d IDPDebt, yearFrom, yearUntil, aumFrom, aumUntil, debtFrom, debtUntil int) int {
	// inclusive sum
	var rowCount = 0

	for _, row := range d.Data {
		var (
			cndYear = isWithIn(yearFrom, yearUntil, row.LoanDate[:4])
			cndAUM  = isWithIn(aumFrom, aumUntil, row.AUMTotal)
			cndDebt = isWithIn(debtFrom, debtUntil, row.LoanAmount)
		)

		if cndYear && cndAUM && cndDebt {
			rowCount += 1
		}
	}

	return rowCount
}
