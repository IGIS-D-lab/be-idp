package v2

import (
	"IGISBackEnd/orm"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/go-redis/redis"
)

func GetMacro(database *redis.Client, w http.ResponseWriter, r *http.Request) {
	var result []IDPMacro
	d, err := orm.JSONGet[[]IDPMacro](database, "macro_asset:1", "$", &result)
	if err != nil {
		fmt.Println(err)
	}

	values := r.URL.Query()
	seg := d[0].search(values)
	packet, _ := json.Marshal(seg)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func PostMacro(database *redis.Client, domestic bool, w http.ResponseWriter, r *http.Request) {
	var newValue newMacroPost
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(body, &newValue)

	switch domestic {
	case true:
		addNewDomestic(database, newValue)
	case false:
		addNewForeign(database, newValue)
	}
}

func addNewDomestic(rdb *redis.Client, n newMacroPost) {
	key := "macro_asset:1"
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.kr1y", &n.KR1Y[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.kr3y", &n.KR3Y[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.kr5y", &n.KR5Y[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.ifd1y", &n.IFB1Y[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.cd91d", &n.CD91D[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.cp91d", &n.CP91D[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.koribor3m", &n.KORIBOR3M[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.fb6m", &n.FB6M[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.fb1y", &n.FB1Y[0])
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.fb3y", &n.FB3Y[0])
}

func addNewForeign(rdb *redis.Client, n newMacroPost) {
	key := "macro_asset:1"
	orm.JSONArrAppend[macroRow](rdb, key, "$.data.ffr", &n.Feds[0])
}

func (d IDPMacro) search(v url.Values) IDPMacro {
	var (
		dateFrom  = v.Get("dateFrom")
		dateUntil = v.Get("dateUntil")
	)

	if (dateFrom == "") && (dateUntil == "") {
		return d
	}

	sendPacket := IDPMacro{
		FactSheet:   d.FactSheet,
		Description: d.Description,
		LastUpdate:  d.LastUpdate,
		Data: macros{
			KR1Y:      d.Data.KR1Y.searchDate(dateFrom, dateUntil),
			KR3Y:      d.Data.KR3Y.searchDate(dateFrom, dateUntil),
			KR5Y:      d.Data.KR5Y.searchDate(dateFrom, dateUntil),
			IFD1Y:     d.Data.IFD1Y.searchDate(dateFrom, dateUntil),
			CD91D:     d.Data.CD91D.searchDate(dateFrom, dateUntil),
			CP91D:     d.Data.CP91D.searchDate(dateFrom, dateUntil),
			KORIBOR3M: d.Data.KORIBOR3M.searchDate(dateFrom, dateUntil),
			Feds:      d.Data.Feds.searchDate(dateFrom, dateUntil),
			FB6M:      d.Data.FB6M.searchDate(dateFrom, dateUntil),
			FB1Y:      d.Data.FB1Y.searchDate(dateFrom, dateUntil),
			FB3Y:      d.Data.FB3Y.searchDate(dateFrom, dateUntil),
		},
	}
	return sendPacket
}

/*
	searchDate for macroRows
	- overwrite slice method by creating another type macroRows
	- condition search date
*/
func (m macroRows) searchDate(dateStart, dateEnd string) macroRows {
	var result macroRows
	for _, row := range m {
		if (row.Date >= dateStart) && (row.Date <= dateEnd) {
			result = append(result, row)
		}
	}
	return result
}
