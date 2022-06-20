package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeAssetWhole(d IDPAsset, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// process parameters
	// s in ["All", "Core", "Value-added", "Opportunistic"]
	s := params["strat"]
	log.Println("Serve asset(checklist) where strategy is ", s)
	qry := ReqIDPAsset{
		Strategy: s,
	}

	packet, _ := json.Marshal(procAssetQry(qry, d))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(packet)
}

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
