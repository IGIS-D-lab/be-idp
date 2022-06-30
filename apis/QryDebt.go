package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

func ServeDebt(d IDPDebt, epType int, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	log.Println(MSG_DEBT, values)

	rows, dt, dg1, dg2 := procDebtQry(values, d, epType)

	sendPacket := IDPDebt{
		FromSheet:      d.FromSheet,
		Desc:           d.Desc,
		LastUpdate:     d.LastUpdate,
		Data:           dt,
		DataGraphLeft:  dg1,
		DataGraphRight: dg2,
		RowCount:       rows,
	}

	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func getFieldDebt(v *debts, field string) string {
	r := reflect.ValueOf(v)
	f := reflect.Indirect(r).FieldByName(field)
	return f.String()
}

func divDebtArray(d []debts, pageNum string) []debts {
	if pageNum == "" {
		return d
	}
	p, _ := strconv.Atoi(pageNum)
	s, e := SINGLE_PAGE_INFO*(p-1), SINGLE_PAGE_INFO*p
	if e > len(d) {
		return d[s:]
	} else {
		return d[s:e]
	}

}

func procDebtQry(v url.Values, d IDPDebt, forGraph int) (int, []debts, []debtsGraphLeft, []debtsGraphRight) {
	var (
		// string parameters - data search
		assetType = v.Get("at")
		seniority = v.Get("seniorstr")
		loanClass = v.Get("loancls")
		rate      = v.Get("rate")
	)
	var (
		// float64 parameters - data search
		debtFrom  = v.Get("debtFrom")
		debtUntil = v.Get("debtUntil")
	)
	_ = map[string]string{
		"at":        assetType,
		"seniorstr": seniority,
		"loancls":   loanClass,
		"rate":      rate,
	} // fore middleware

	// TODO: create sorting
	var (
		_   = v.Get("sortOrd")
		_   = v.Get("sorkKey")
		pgn = v.Get("pageCount")
	)
	var (
		dataPoints = 0
		// conditions
		sendPacketDT = []debts{}
		sendPacketG1 = []debtsGraphLeft{}
		sendPacketG2 = []debtsGraphRight{}
	)

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
				dataPoints += 1

			case 1: // graph left
				r := debtsGraphLeft{
					SetDateRate: row.SetDateRate,
					AssetType:   row.AssetType,
					AssetName:   row.AssetName,
					LoanDate:    row.LoanDate,
					LoanAmount:  row.LoanAmount,
				}
				sendPacketG1 = append(sendPacketG1, r)
				dataPoints += 1

			case 2: // graph right
				r := debtsGraphRight{
					LoanDate:   row.LoanDate,
					LoanAmount: row.LoanAmount,
					LPCorp:     row.LPCorp,
				}
				sendPacketG2 = append(sendPacketG2, r)
				dataPoints += 1
			}
		}
	}
	return dataPoints, divDebtArray(sendPacketDT, pgn), sendPacketG1, sendPacketG2
}
