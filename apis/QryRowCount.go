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
	// process parameters
	yearFrom, err := strconv.Atoi(params["yearFrom"])
	if err != nil {
		log.Println("QryRowCount :: yearFrom not integer")
	}
	yearUntil, err := strconv.Atoi(params["yearUntil"])
	if err != nil {
		log.Println("QryRowCount :: yearUntil not integer")
	}
	err = verifyRowCountYear(yearFrom, yearUntil)
	aumFrom, err := strconv.ParseFloat(params["aumFrom"], 64)
	if err != nil {
		log.Println("QryRowCount :: aumFrom not float64", params["aumFrom"])
	}
	aumUntil, err := strconv.ParseFloat(params["aumUntil"], 64)
	if err != nil {
		log.Println("QryRowCount :: aumUntil not float64", params["aumUntil"])
	}
	debtFrom, err := strconv.ParseFloat(params["debtFrom"], 64)
	if err != nil {
		log.Println("QryRowCount :: debtFrom not float64")
	}
	debtUntil, err := strconv.ParseFloat(params["debtUntil"], 64)
	if err != nil {
		log.Println("QryRowCount :: debtUntil not float64")
	}
	qry := ReqRowCount{
		YearFrom:  yearFrom,
		YearUntil: yearUntil,
		AumFrom:   aumFrom,
		AumUntil:  aumUntil,
		DebtFrom:  debtFrom,
		DebtUntil: debtUntil,
	}

	sendPacket := IDPRowCount{
		RowCount: procRowCountQry(d, qry),
	}
	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func verifyRowCountYear(yearFrom, yearUntil int) error {
	if yearFrom <= yearUntil {
		return nil
	} else {
		return errors.New("QryRowCount :: yearFrom is bigger than yearUntil")
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

func procRowCountQry(d IDPDebt, rq ReqRowCount) int {
	// inclusive sum
	var rowCount = 0
	aumFromInt, aumUntilInt := int(rq.AumFrom), int(rq.AumUntil)
	debtFromInt, debtUntilInt := int(rq.DebtFrom), int(rq.DebtUntil)
	for _, row := range d.Data {
		var (
			cndYear = isWithIn(rq.YearFrom, rq.YearUntil, row.LoanDate[:4])
			cndAUM  = isWithIn(aumFromInt, aumUntilInt, row.AUMTotal)
			cndDebt = isWithIn(debtFromInt, debtUntilInt, row.LoanAmount)
		)

		if cndYear && cndAUM && cndDebt {
			rowCount += 1
		}
	}

	return rowCount
}
