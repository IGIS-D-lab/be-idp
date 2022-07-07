package apis

import (
	"encoding/json"
	"net/http"
	"net/url"
)

/*
	ServeModelInfo
	- Serve idpModelInfo.json file directly without any queries
*/
func ServeModelInfo(d IDPModelInfo, w http.ResponseWriter, r *http.Request) {
	packet, _ := json.Marshal(d)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

/*
	ServeModelCoef
	- Serve idpModelCoef.json file directly without any queries
*/
func ServeModelCoef(d IDPModelCoef, w http.ResponseWriter, r *http.Request) {
	packet, _ := json.Marshal(d)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

/*
	findRecentMacro
	- from macro data rows, compare dates (YYYYMMDD str)
	- return macro data row containing maximum dates
*/
func findRecentMacro(d []macroRow) macroRow {
	var (
		maxDate = "00000000" // initial state
		recent  = macroRow{}
	)
	for _, r := range d {
		if r.Date >= maxDate {
			recent = r
		}
	}
	return recent
}

/*
	genDataPointMap
	- Coefficient 'b' in the model is mapped with integer
	  - -1: constant etc.
	- Function that makes 'X'
	- Provide map[ coefficient key ] = datapoint value
*/
func genDataPointMap(mac IDPMacro) map[int]float64 {
	// get recent struct rows within O(n) time - TODO: change it with database
	var result = map[int]float64{}
	result[-1] = 1
	result[0] = findRecentMacro(mac.Data.KR1Y).Value
	result[1] = findRecentMacro(mac.Data.KR3Y).Value
	result[2] = findRecentMacro(mac.Data.KR5Y).Value
	result[3] = findRecentMacro(mac.Data.IFD1Y).Value
	result[4] = findRecentMacro(mac.Data.CD91D).Value
	result[5] = findRecentMacro(mac.Data.CP91D).Value
	result[6] = findRecentMacro(mac.Data.KORIBOR3M).Value
	return result
}

/*
	genParameterMap
	- Function that makes 'b'
	- map[ coefficient key ] = coefficient
*/
func genParameterMap(coefs []coefficient) map[int]float64 {
	// coef json file. - each model variable has an integer designated to it
	// match that integer with coefficient.
	var result = map[int]float64{}
	for _, c := range coefs {
		result[c.CoefIndex] = c.Coef
	}
	return result
}

/*
	procModeQuery
	- returns user searched
	- 1) loanClassMap - model parameter integer key for loancls
	- 2) seniorityMap - model parameter integer key for seniorstr
*/
func procModelQuery(v url.Values) (int, int) {
	var (
		seniority = v.Get("seniorstr")
		loanClass = v.Get("loancls")
	)
	var (
		loanClassMap = map[string]int{
			"PF":  10,
			"담보":  11,
			"부가세": 12,
			"브릿지": 13,
			"한도":  14,
		}
		seniorityMap = map[string]int{
			"선": 7,
			"중": 8,
			"후": 9,
		}
	)
	return loanClassMap[loanClass], seniorityMap[seniority]
}

/*
	calcInterest
	- get 2 map[int]float x, b
	- x [ coefficient key ] * b [ coefficient key ]

	- return per liqProvider and intRate
*/
func calcInterest(x, b map[int]float64, liqProv, intRate int) float64 {
	var res float64
	for multiKey, val := range b {
		if (multiKey == liqProv) || (multiKey == intRate) {
			res += val
		} else {
			res += (x[multiKey] * b[multiKey])
		}
	}
	return res
}

/*
	ServeModelCalc
	- Fill in ModelPrediction
	- liquidity provider in Bank, Insurance(Ins), Etc
	- rate in Fix, Float
*/
func ServeModelCalc(model IDPModelCoef, band IDPModelInfo, macro IDPMacro, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()

	// set x
	// x: categorical data
	userSearch1, userSearch2 := procModelQuery(values)
	// x: recent Macro Values
	x := genDataPointMap(macro)
	x[userSearch1] = 1
	x[userSearch2] = 1

	// set coefficient
	b := genParameterMap(model.Data)

	sendpacket := ModelPrediction{
		BankFix: []float64{calcInterest(x, b, 19, 15), calcInterest(x, b, 19, 15)},
		InsFix:  []float64{calcInterest(x, b, 18, 15), calcInterest(x, b, 18, 15)},
		EtcFix:  []float64{calcInterest(x, b, 17, 15), calcInterest(x, b, 17, 15)},

		BankFloat: []float64{calcInterest(x, b, 19, 16), calcInterest(x, b, 19, 16)},
		InsFloat:  []float64{calcInterest(x, b, 18, 16), calcInterest(x, b, 18, 16)},
		EtcFloat:  []float64{calcInterest(x, b, 17, 16), calcInterest(x, b, 17, 16)},
	}
	packet, _ := json.Marshal(sendpacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}
