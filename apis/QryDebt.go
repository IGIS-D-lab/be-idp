package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func ServeDebt(d IDPDebt, epType int, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	log.Println(MSG_DEBT, values)

	dt, dg1, dg2 := procDebtQry(values, d, epType)

	sendPacket := IDPDebt{
		FromSheet:      d.FromSheet,
		Desc:           d.Desc,
		LastUpdate:     d.LastUpdate,
		Data:           dt,
		DataGraphLeft:  dg1,
		DataGraphRight: dg2,
	}
	sendPacket.RowCount = len(sendPacket.Data)

	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func procDebtQry(v url.Values, d IDPDebt, forGraph int) ([]debts, []debtsGraphLeft, []debtsGraphRight) {

	// Queries
	var (
		assetType = v.Get("at")
		seniority = v.Get("seniorstr")
		loanClass = v.Get("loancls")
		debtFrom  = v.Get("debtFrom")
		debtUntil = v.Get("debtUntil")
		rate      = v.Get("rate")
	)

	// Conditions
	var sendPacketDT = []debts{}
	var sendPacketG1 = []debtsGraphLeft{}
	var sendPacketG2 = []debtsGraphRight{}

	for _, row := range d.Data {
		var (
			cndAssetType  = true
			cndSeniorty   = true
			cndLoanClass  = true
			cndRate       = true
			cndDebtAmount = true
		)

		if assetType != "" {
			cndAssetType = IsWithInChoice(assetType, row.AssetType)
		}
		if seniority != "" {
			cndSeniorty = IsWithInChoice(seniority, row.Seniority)
		}
		if rate != "" {
			cndRate = IsWithInChoice(rate, row.RateType)
		}
		if loanClass != "" {
			cndLoanClass = IsWithInChoice(loanClass, row.LoanClass)
		}
		if (debtFrom != "") && (debtUntil != "") {
			df, _ := strconv.ParseFloat(debtFrom, 64)
			du, _ := strconv.ParseFloat(debtUntil, 64)
			cndDebtAmount = IsWithInSlider(int(df), int(du), row.LoanAmount)
		}

		if cndAssetType && cndSeniorty && cndLoanClass && cndRate && cndDebtAmount {
			switch forGraph {
			case 0: // table
				sendPacketDT = append(sendPacketDT, row)
			case 1: // graph left
				r := debtsGraphLeft{
					SetDateRate: row.SetDateRate,
					AssetType:   row.AssetType,
					LoanDate:    row.LoanDate,
					LoanAmount:  row.LoanAmount,
				}
				sendPacketG1 = append(sendPacketG1, r)
			case 2: // graph right
				r := debtsGraphRight{
					LoanDate:   row.LoanDate,
					LoanAmount: row.LoanAmount,
					LPCorp:     row.LPCorp,
				}
				sendPacketG2 = append(sendPacketG2, r)
			}
		}
	}
	return sendPacketDT, sendPacketG1, sendPacketG2
}
