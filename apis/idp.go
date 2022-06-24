package apis

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const SINGLE_PAGE_INFO = 10

const (
	DATA_PANIC_ASSET = "Checklist :: Panic :: "
	DATA_PANIC_DEBT  = "Debt :: Panic :: "
	DATA_PANIC_MACRO = "Macro :: Panic :: "

	DATA_ERR_ASSET = "Checklist :: Error :: "
	DATA_ERR_DEBT  = "Debt :: Error :: "
	DATA_ERR_MACRO = "Macro :: Error :: "
)

const (
	MSG_ASSET  = "Asset :: "
	MSG_DEBT   = "Debt :: "
	MSG_MACRO  = "Macro :: "
	MSG_SINGLE = "Single :: "
)

const (
	TEST_URL_ASSET  = "http://localhost:8080/api/v1/asset?strat=Core"
	TEST_URL_DEBT   = "http://localhost:8080/api/v1/debt/dataTable?at=%EC%98%A4%ED%94%BC%EC%8A%A4-%ED%98%B8%ED%85%94&seniorstr=%EC%84%A0&loancls=%EB%B8%8C%EB%A6%BF%EC%A7%80&debtFrom=1&debtUntil=1e13&pageCount=1"
	TEST_URL_MACRO  = "http://localhost:8080/api/v1/macro?commodity=kr1y&yearFrom=2010&yearUntil=2020"
	TEST_URL_SINGLE = "http://localhost:8080/api/v1/single?fc=112001&idx=1"
)

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IGIS Debt Platform landing page\n")
}

func IsWithInSlider(s, e int, val string) bool {
	v := strings.Replace(val, ",", "", -1)
	res, _ := strconv.Atoi(v)
	if (res >= s) && (res <= e) {
		return true
	} else {
		return false
	}
}

func IsWithInChoice(isSame, val string) bool {
	workString := strings.Replace(isSame, " ", "", -1)
	workStringSlice := strings.Split(workString, "-")
	result := false
	for _, i := range workStringSlice {
		if i == val {
			result = true
		}
	}
	return result
}
