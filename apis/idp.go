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
	DATA_PANIC_MODEL = "Model :: Panic :: "

	DATA_ERR_ASSET = "Checklist :: Error :: "
	DATA_ERR_DEBT  = "Debt :: Error :: "
	DATA_ERR_MACRO = "Macro :: Error :: "
	DATA_ERR_MODEL = "Model :: Error :: "
)

const (
	MSG_ASSET  = "Asset :: "
	MSG_DEBT   = "Debt :: "
	MSG_MACRO  = "Macro :: "
	MSG_SINGLE = "Single :: "
	MSG_MODEL  = "Model :: "
)

/*
	ServeLanding
	- landing page
*/
func ServeLanding(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IGIS Debt Platform landing page\n")
}

func IsWithInDate(s, e string, val string) bool {
	if (val >= s) && (val <= e) {
		return true
	} else {
		return false
	}
}

/*
	IsWithInSlider
	- return true if val is within slider s and e
*/
func IsWithInSlider(s, e int, val string) bool {
	v := strings.Replace(val, ",", "", -1)
	res, _ := strconv.Atoi(v)
	if (res >= s) && (res <= e) {
		return true
	} else {
		return false
	}
}

/*
	IsWithInChoice
	- return true if val is within multiple choice
	- multiple choice is connected by "-"
*/
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
