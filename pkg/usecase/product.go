package usecase

import (
	"github.com/kben19/CleoApp/pkg/domain"
	"github.com/kben19/CleoApp/pkg/lib/common"
	"github.com/kben19/CleoApp/pkg/types"
)

type UsecaseProduct struct {
	domain domain.DomainProductItf
}

type UsecaseProductItf interface {
	GetProductByID(id int) (types.Product, error)
}

func InitUsecaseProduct(domain domain.DomainProductItf) UsecaseProduct {
	return UsecaseProduct{domain: domain}
}

func (u UsecaseProduct) GetProductByID(id int) (types.Product, error) {
	if id <= 0 {
		return types.Product{}, common.ErrInvalidIDParameter
	}

	product, err := u.domain.GetProductByID(id)
	if err != nil {
		return product, err
	}
	return product, nil
}
