package server

import (
	"database/sql"

	"github.com/kben19/CleoApp/pkg/domain"
	"github.com/kben19/CleoApp/pkg/usecase"
)

type usecases struct {
	product usecase.UsecaseProductItf
}

// Init function from usecases to resource
func Init(db *sql.DB) usecases {
	useCase := usecases{}

	productDomain := domain.InitDomainProduct(db)

	useCase.product = usecase.InitUsecaseProduct(productDomain)

	return useCase
}
