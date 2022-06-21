package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func ServeDebt(d IDPDebt, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	log.Println(MSG_DEBT, values)

	sendPacket := IDPDebt{
		FromSheet:  d.FromSheet,
		Desc:       d.Desc,
		LastUpdate: d.Desc,
		Data:       procDebtQry(values, d),
	}
	sendPacket.RowCount = len(sendPacket.Data)

	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func procDebtQry(v url.Values, d IDPDebt) []debts {
	// Queries
	var (
		assetType = v.Get("at")
		seniority = v.Get("seniorstr")
		loanClass = v.Get("loancls")
		debtFrom  = v.Get("debtFrom")
		debtUntil = v.Get("debtUntil")
	)

	// Conditions
	var sendPacket = []debts{}

	for _, row := range d.Data {
		var (
			cndAssetType  = true
			cndSeniorty   = true
			cndLoanClass  = true
			cndDebtAmount = true
		)

		if assetType != "" {
			cndAssetType = IsWithInChoice(assetType, row.AssetType)
		}
		if seniority != "" {
			cndSeniorty = IsWithInChoice(seniority, row.Seniority)
		}
		if loanClass != "" {
			cndLoanClass = IsWithInChoice(loanClass, row.LoanClass)
		}
		if (debtFrom != "") && (debtUntil != "") {
			df, _ := strconv.ParseFloat(debtFrom, 64)
			du, _ := strconv.ParseFloat(debtUntil, 64)
			cndDebtAmount = IsWithInSlider(int(df), int(du), row.LoanAmount)
		}
		if cndAssetType && cndSeniorty && cndLoanClass && cndDebtAmount {
			sendPacket = append(sendPacket, row)
		}
	}
	return sendPacket
}
