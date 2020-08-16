package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type HTTPMetaResponse struct {
	StatusCode int         `json:"status_code"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
	Param      interface{} `json:"param"`
}

// HTTPError extension function of http.Error that incorporate json meta
func HTTPError(w http.ResponseWriter, err error, msg string, code int, param ...interface{}) {
	if code <= 0 {
		code = http.StatusInternalServerError
	}
	if err == nil {
		code = http.StatusOK
		err = errors.New("")
	}
	if msg == "" {
		msg = http.StatusText(code)
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	raw, errMarshal := json.Marshal(HTTPMetaResponse{
		StatusCode: code,
		Error:      err.Error(),
		Message:    msg,
		Param:      param,
	})
	if errMarshal != nil {
		http.Error(w, errMarshal.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintln(w, string(raw))
}
