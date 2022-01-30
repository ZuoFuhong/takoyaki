package web

import (
	"encoding/json"
	"io"
	"net/http"
	"takoyaki/defs"
)

func SendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(errResp.HttpSC)
	resStr, _ := json.Marshal(&errResp.Error)
	_, _ = io.WriteString(w, string(resStr))
}

func SendNormalJsonResponse(w http.ResponseWriter, body interface{}) {
	bytes, _ := json.Marshal(body)
	SendNormalResponse(w, string(bytes), http.StatusOK)
}

func SendNormalResponse(w http.ResponseWriter, body string, sc int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(sc)
	_, _ = io.WriteString(w, body)
}
