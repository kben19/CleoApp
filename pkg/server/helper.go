package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kben19/CleoApp/pkg/lib/common"
	"github.com/kben19/CleoApp/pkg/types"
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
	w.Header().Set("Content-Type", "application/json")
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

func WriteHTTPResponse(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	raw, err := json.Marshal(res)
	if err != nil {
		HTTPError(w, err, common.FailedToMarshal, http.StatusInternalServerError, res)
		return
	}
	_, err = w.Write(raw)
	if err != nil {
		HTTPError(w, err, common.FailedToWriteResponse, http.StatusInternalServerError)
		return
	}
	return
}

func GetListParam(w http.ResponseWriter, r *http.Request) (types.StandardListParam, error) {
	limit, err := strconv.Atoi(r.FormValue(common.LimitParam))
	if err != nil {
		HTTPError(w, err, common.InvalidLimitParameterMsg, http.StatusBadRequest)
		return types.StandardListParam{}, err
	}
	offset, err := strconv.Atoi(r.FormValue(common.OffsetParam))
	if err != nil {
		HTTPError(w, err, common.InvalidOffsetParameterMsg, http.StatusBadRequest)
		return types.StandardListParam{}, err
	}

	return types.StandardListParam{
		Offset:    offset,
		Limit:     limit,
		OrderBy:   r.FormValue(common.OrderByParam),
		OrderType: r.FormValue(common.OrderTypeParam),
	}, nil
}
