package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

/*
	ServeSingle
	- serve need to provide detail information
	- single out fund by its unique index 'idx'
*/
func ServeSingle(d IDPDebt, w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	log.Println(MSG_SINGLE, values)

	dp := procSingleQry(values, d)

	sendPacket := IDPSingle{
		FromSheet:  d.FromSheet,
		Desc:       d.Desc,
		LastUpdate: d.LastUpdate,
		Data:       dp,
	}
	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func procSingleQry(v url.Values, d IDPDebt) debts {
	var (
		fundCode = v.Get("fc")
		uniqIdx  = v.Get("idx")
	)
	var sendPacketSingle = debts{}

	for _, row := range d.Data {
		sameFundCode := row.FundCode == fundCode
		sameLP := row.UniqueIndex == uniqIdx
		if sameFundCode && sameLP {
			sendPacketSingle = row
			break
		}
	}
	return sendPacketSingle
}
