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
		log.Panicln("Checklist :: Panic :: ", err)
	}

	debtData, err := mntDebt()
	if err != nil {
		log.Panicln("Debt :: Panic :: ", err)
	}

	macroData, err := mntMacro()
	if err != nil {
		log.Panicln("Macro :: Panic :: ", err)
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
		log.Println("Checklist :: Err :: ", err)
	} else {
		log.Println("Checklist :: Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPAsset
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println("Asset :: Err :: ", err)
		return data, err
	} else {
		return data, nil
	}
}

func mntDebt() (IDPDebt, error) {
	file, err := os.Open("./asset/idpDebt.json")
	if err != nil {
		log.Println("Debt :: Err :: ", err)
	} else {
		log.Println("Debt :: Successfully opened")
	}

	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	var data IDPDebt
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println("Debt :: Err :: ", err)
		return data, err
	} else {
		return data, nil
	}
}

func mntMacro() (IDPMacro, error) {
	file, err := os.Open("./asset/idpMacro.json")
	if err != nil {
		log.Println("Macro :: Err :: ", err)
	} else {
		log.Println("Macro :: Successfully opened")
	}
	byteVal, _ := ioutil.ReadAll(file)
	byteVal = bytes.Replace(byteVal, []byte(": NaN"), []byte(":null"), -1)

	// JSON Marshalling
	var data IDPMacro
	err = json.Unmarshal(byteVal, &data)
	if err != nil {
		log.Println("Macro :: Err :: ", err)
		return data, err
	} else {
		return data, nil
	}

}
