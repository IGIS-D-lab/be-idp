package apis

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

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

	return IDPDataSet{
		Asset: assetData,
		Debt:  debtData,
		Macro: macroData,
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
	file, err := os.Open("./asset/idpDebt.json")
	if err != nil {
		log.Println(DATA_ERR_DEBT, err)
	} else {
		log.Println(MSG_DEBT, "Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPDebt
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println(DATA_ERR_DEBT, err)
		return data, err
	} else {
		return data, nil
	}
}

func mntMacro() (IDPMacro, error) {
	file, err := os.Open("./asset/idpMacro.json")
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
