package server

import (
	"net/http"
	"strconv"

	"github.com/kben19/CleoApp/pkg/lib/common"
	"github.com/kben19/CleoApp/pkg/usecase"
)

const (
	failedGetProductMsg     = "Failed to get product"
	failedGetProductListMsg = "Failed to get product list"
)

type HandlerProduct struct {
	productUsecase usecase.UsecaseProductItf
}

func InitHandlerProduct(productUsecase usecase.UsecaseProductItf) HandlerProduct {
	return HandlerProduct{productUsecase: productUsecase}
}

func (h HandlerProduct) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.FormValue(common.IDParam))
	if err != nil {
		HTTPError(w, err, common.InvalidIDParameterMsg, http.StatusBadRequest)
		return
	}

	res, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		HTTPError(w, err, failedGetProductMsg, http.StatusInternalServerError, id)
		return
	}
	WriteHTTPResponse(w, res)
	return
}

func (h HandlerProduct) handleGetProductList(w http.ResponseWriter, r *http.Request) {
	param, err := GetListParam(w, r)
	if err != nil {
		return
	}

	err = param.ValidateProductParam()
	if err != nil {
		HTTPError(w, err, common.InvalidParamaterMsg, http.StatusBadRequest)
		return
	}

	products, err := h.productUsecase.GetProductListWithPagination(param)
	if err != nil {
		HTTPError(w, err, failedGetProductListMsg, http.StatusInternalServerError)
		return
	}
	WriteHTTPResponse(w, products)
	return
}
