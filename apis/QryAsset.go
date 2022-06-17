package apis

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ServeAssetWhole(d IDPAsset, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// process strategy s ["Core", "Value-added", "Opportunistic"]
	s := params["strat"]
	log.Println("Serve asset(checklist) where strategy is ", s)

	if s == "All" {
		// push all strategy
		packet, _ := json.Marshal(d.Data)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	} else {
		// push specific strategy
		var sendPacket = []assets{}
		for _, row := range d.Data {
			switch {
			case s == row.Strategy:
				sendPacket = append(sendPacket, row)
			default:
			}
		}

		packet, _ := json.Marshal(sendPacket)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}

}
