package logs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type RequestAmount struct {
	AssetTypeReq map[string]int `json:"Qat"`
	SeniorityReq map[string]int `json:"Qseniorstr"`
	LoanClassReq map[string]int `json:"Qloancls"`
	RateReq      map[string]int `json:"Qrate"`
}

type LogRequestAmount struct {
	User string        `json:"user"`
	Data RequestAmount `json:"data"`
}

func createLog() {

}

func LogInit() (LogRequestAmount, error) {
	var data LogRequestAmount

	// open data
	file, err := os.Open("./logs/log.json")
	if err != nil {
		log.Println(err)
		return data, err
	} else {
		log.Println(file)
	}
	defer file.Close()

	// json reading
	byteVal, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println(err)
		return data, err
	} else {
		return data, nil
	}
}
