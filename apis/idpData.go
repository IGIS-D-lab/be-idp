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

func parseResult(redisJson string) []macroRow {
	var redisRows [][]macroRow
	byteVal := []byte(redisJson)
	_ = json.Unmarshal(byteVal, &redisRows)
	return redisRows[0]
}

func mntMacroRedis(db *redis.Client) IDPMacro {
	kr1y, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.kr1y").Result()
	kr3y, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.kr3y").Result()
	kr5y, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.kr5y").Result()
	ifd1y, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.ifd1y").Result()
	cd91d, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.cd91d").Result()
	cp91d, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.cp91d").Result()
	koribor3m, _ := orm.RedisJSONGet(db, "macro_asset:1", "$.data.koribor3m").Result()

	return IDPMacro{
		Data: macros{
			KR1Y:      parseResult(kr1y),
			KR3Y:      parseResult(kr3y),
			KR5Y:      parseResult(kr5y),
			IFD1Y:     parseResult(ifd1y),
			CD91D:     parseResult(cd91d),
			CP91D:     parseResult(cp91d),
			KORIBOR3M: parseResult(koribor3m),
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
