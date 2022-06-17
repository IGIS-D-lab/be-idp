package apis

import (
	"fmt"
	"net/http"
)

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IGIS Debt Platform landing page\n")
}
