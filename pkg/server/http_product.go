package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kben19/CleoApp/pkg/lib/common"
	"github.com/kben19/CleoApp/pkg/usecase"
)

const (
	failedGetProductMsg = "Failed to get product"
)

type HandlerProduct struct {
	productUsecase usecase.UsecaseProductItf
}

func InitHandlerProduct(productUsecase usecase.UsecaseProductItf) HandlerProduct {
	return HandlerProduct{productUsecase: productUsecase}
}

func (h HandlerProduct) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue(common.IDParams))
	if err != nil {
		HTTPError(w, err, common.InvalidIDParameterMsg, http.StatusBadRequest)
		return
	}

	res, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		HTTPError(w, err, failedGetProductMsg, http.StatusInternalServerError, id)
		return
	}

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
