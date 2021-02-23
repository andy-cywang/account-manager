package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	RequestFailed = "failed"
)

type Result struct {
	Body struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
	}
}

func WriteJSONResponse(w http.ResponseWriter, obj interface{}, httpCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}

func WriteErrResponse(w http.ResponseWriter, err error, httpCode int) {
	r := Result{}
	r.Body.Status = RequestFailed
	r.Body.Message = err.Error()

	log.Println("middleware: " + strings.ToLower(http.StatusText(httpCode)) + " - " + err.Error())
	WriteJSONResponse(w, r.Body, httpCode)
}