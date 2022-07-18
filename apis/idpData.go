package apis

import (
	"IGISBackEnd/orm"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-redis/redis"
)

/*
	MntData
	- mount data before you serve API
*/
func MntData() IDPDataSet {
	assetData, err := mntAsset()
	if err != nil {
		log.Panicln(DATA_PANIC_ASSET, err)
	}

	debtData, err := mntDebt()
	if err != nil {
		log.Panicln(DATA_PANIC_DEBT, err)
	}

	macroData, err := mntMacro()
	if err != nil {
		log.Panicln(DATA_PANIC_MACRO, err)
	}

	modelInfo, err := mntModelInfo()
	if err != nil {
		log.Panicln(DATA_PANIC_MODEL, err)
	}
	modelCoef, err := mntModelCoef()
	if err != nil {
		log.Panicln(DATA_PANIC_MODEL, err)
	}

	return IDPDataSet{
		Asset:     assetData,
		Debt:      debtData,
		Macro:     macroData,
		ModelInfo: modelInfo,
		ModelCoef: modelCoef,
	}
}

func mntAsset() (IDPAsset, error) {
	file, err := os.Open("./asset/idpChecklist.json")
	if err != nil {
		log.Println(DATA_ERR_ASSET, err)
	} else {
		log.Println(MSG_ASSET, "Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPAsset
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println(DATA_ERR_ASSET, err)
		return data, err
	} else {
		return data, nil
	}
}

func mntDebt() (IDPDebt, error) {
	file, err := os.Open("./asset/idpDebt3.json")
	if err != nil {
		log.Println(DATA_ERR_DEBT, err)
	} else {
		log.Println(MSG_DEBT, "Successfully opened")
	}

	// python JSON compiler gives : NaN
	// go cannot read this. Therefore changing it to :null
	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPDebt
	err = json.Unmarshal(byteVal, &data)
	return data, err
}

func mntMacro() (IDPMacro, error) {
	file, err := os.Open("./asset/idpMacro4.json")
	if err != nil {
		log.Println(DATA_ERR_MACRO, err)
	} else {
		log.Println(MSG_MACRO, "Successfully opened")
	}
	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	// JSON Marshalling
	var data IDPMacro
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println(DATA_ERR_MACRO, err)
		return data, err
	} else {
		return data, nil
	}

}

func mntMacroRedis(db *redis.Client) IDPMacro {
	var rContain []macros

	data, err := orm.JSONGet[[]macros](db, "macro_asset:1", "$.data", &rContain)
	if err != nil {
		return IDPMacro{}
	}
	return IDPMacro{
		Data: macros{
			KR1Y:      data[0].KR1Y,
			KR3Y:      data[0].KR3Y,
			KR5Y:      data[0].KR5Y,
			IFD1Y:     data[0].IFD1Y,
			CD91D:     data[0].CD91D,
			CP91D:     data[0].CP91D,
			KORIBOR3M: data[0].KORIBOR3M,
		},
	}

}

func mntModelInfo() (IDPModelInfo, error) {
	file, err := os.Open("./asset/idpModelInfo.json")
	if err != nil {
		log.Println(DATA_ERR_MODEL, err)
	} else {
		log.Println(MSG_MODEL, "Successfully opened")
	}
	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPModelInfo
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println(DATA_ERR_MODEL, err)
		return data, err
	} else {
		return data, nil
	}
}

func mntModelCoef() (IDPModelCoef, error) {
	file, err := os.Open("./asset/idpCoef2.json")
	if err != nil {
		log.Println(DATA_ERR_MODEL, err)
	} else {
		log.Println(MSG_MODEL, "Successfully opened")
	}
	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPModelCoef
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println(DATA_ERR_MODEL, err)
		return data, err
	} else {
		return data, nil
	}
}
