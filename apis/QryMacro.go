package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ServeMacroWhole(d IDPMacro, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// process year slider
	yearFrom, err := strconv.Atoi(params["yearFrom"])
	if err != nil {
		log.Println("QryDebt :: yearFrom not integer")
	}
	yearUntil, err := strconv.Atoi(params["yearUntil"])
	if err != nil {
		log.Println("QryDebt :: yearUntil not integer")
	}
	searchComm := params["commodity"]

	var sendPacket = []macroRow{}
	switch {
	case searchComm == "kr1y":
		for _, row := range d.Data.KR1Y {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	case searchComm == "kr3y":
		for _, row := range d.Data.KR3Y {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	case searchComm == "kr5y":
		for _, row := range d.Data.KR5Y {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	case searchComm == "ifd1y":
		for _, row := range d.Data.IFD1Y {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	case searchComm == "cd91d":
		for _, row := range d.Data.CD91D {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	case searchComm == "cp91d":
		for _, row := range d.Data.CP91D {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	case searchComm == "koribor3m":
		for _, row := range d.Data.KORIBOR3M {
			y, _ := strconv.Atoi(row.Date[:4])
			if (y >= yearFrom) && (y <= yearUntil) {
				sendPacket = append(sendPacket, row)
			}
		}
	}

	packet, _ := json.Marshal(sendPacket)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	// mount json
	w.Write(packet)
}
