package util

import (
	"net/http"
)

func RespText(w http.ResponseWriter, ret string) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.Write([]byte(ret))
}
