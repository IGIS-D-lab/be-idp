package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

/*
	ServeMacro
	- serve dataTable aker
*/
func ServeMacro(d IDPMacro, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	log.Println(MSG_MODEL, values)

	dt := procMacroQuery(values, d)
	packet, _ := json.Marshal(dt)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func procMacroQuery(v url.Values, d IDPMacro) IDPMacro {
	var (
		dateFrom  = v.Get("dateFrom")
		dateUntil = v.Get("dateUntil")
	)

	if (dateFrom == "") && (dateUntil == "") {
		return d
	}

	sendPacket := IDPMacro{
		FromSheet:  d.FromSheet,
		Desc:       d.Desc,
		LastUpdate: d.LastUpdate,
		Data: macros{
			KR1Y:      procByAsset(d.Data.KR1Y, dateFrom, dateUntil),
			KR3Y:      procByAsset(d.Data.KR3Y, dateFrom, dateUntil),
			KR5Y:      procByAsset(d.Data.KR5Y, dateFrom, dateUntil),
			IFD1Y:     procByAsset(d.Data.IFD1Y, dateFrom, dateUntil),
			CD91D:     procByAsset(d.Data.CD91D, dateFrom, dateUntil),
			CP91D:     procByAsset(d.Data.CP91D, dateFrom, dateUntil),
			KORIBOR3M: procByAsset(d.Data.KORIBOR3M, dateFrom, dateUntil),
		},
	}
	return sendPacket
}

func procByAsset(multiRows []macroRow, dtF, dtU string) []macroRow {
	var result = []macroRow{}
	for _, row := range multiRows {
		if IsWithInDate(dtF, dtU, row.Date) {
			result = append(result, row)
		}
	}
	return result
}
