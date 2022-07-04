package apis

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

/*
	ServeAssetWhole
	- API Query processing with Vars
	- TODO: Change it to r.URL.Query - this enables optional existance query
*/
func ServeAssetWhole(d IDPAsset, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	qry := procAssetParam(params)

	packet, _ := json.Marshal(procAssetQry(qry, d))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}

/*
	procAssetParam
	- Store queries in ReqIDPAsset
*/
func procAssetParam(m map[string]string) ReqIDPAsset {
	// strat in ["All", "Core", "Value-added", "Opportunistic"]
	q := ReqIDPAsset{
		Strategy: m["strat"],
	}
	return q
}

/*
	procAssetQry
	- find row inside data that matches condition
*/
func procAssetQry(rq ReqIDPAsset, d IDPAsset) []assets {
	// reads asset query from ReqIDPAsset struct
	var sendPacket = []assets{}
	if rq.Strategy == "All" {
		return d.Data
	}

	for _, row := range d.Data {
		switch {
		case row.Strategy == rq.Strategy:
			sendPacket = append(sendPacket, row)
		default:
		}
	}
	return sendPacket
}
