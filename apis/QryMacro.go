package apis

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"IGISBackEnd/orm"

	"github.com/go-redis/redis"
)

/*
	ServeMacro
	- serve dataTable aker
*/
func ServeMacro(database *redis.Client, w http.ResponseWriter, r *http.Request) {
	d := mntMacroRedis(database)
	values := r.URL.Query()
	log.Println(MSG_MACRO, values)

	dt := procMacroQuery(values, d)
	packet, _ := json.Marshal(dt)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func UpdateMacro(d *IDPMacro, database *redis.Client, w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(MSG_MACRO, err)
	}
	// overwrite memory
	err = procMacroUpdate(body, d)
	if err != nil {
		log.Println("macro overwrite error", err)
	}
	// file updated - reload file
	upd, _ := json.Marshal(d)
	// on memory
	err = ioutil.WriteFile("./asset/idpMacro4.json", upd, 0644)
	// on database
	err = orm.JSONSet[IDPMacro](database, "macro_asset:1", "$", d)
	if err != nil {
		log.Println(err)
	}
	log.Print("update complete")
}

func procMacroUpdate(newVal []byte, d *IDPMacro) error {
	var newMacroRow newMacroPost
	err := json.Unmarshal(newVal, &newMacroRow)
	if err != nil {
		return err
	}
	d.Data.KR1Y = append(d.Data.KR1Y, newMacroRow.KR1Y...)
	d.Data.KR3Y = append(d.Data.KR3Y, newMacroRow.KR3Y...)
	d.Data.KR5Y = append(d.Data.KR5Y, newMacroRow.KR5Y...)
	d.Data.IFD1Y = append(d.Data.IFD1Y, newMacroRow.IFB1Y...)
	d.Data.CD91D = append(d.Data.CD91D, newMacroRow.CD91D...)
	d.Data.CP91D = append(d.Data.CP91D, newMacroRow.CP91D...)
	d.Data.KORIBOR3M = append(d.Data.KORIBOR3M, newMacroRow.KORIBOR3M...)
	return nil
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
			Feds:      procByAsset(d.Data.Feds, dateFrom, dateUntil),
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
