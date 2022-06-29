package apis

import (
	"net/http"
)

func ServeModelInfo(d []byte, w http.ResponseWriter, r *http.Request) {
	w.Write(d)
}

func ServeModelCoef(d []byte, w http.ResponseWriter, r *http.Request) {
	w.Write(d)
}
