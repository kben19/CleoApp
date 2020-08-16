package usecase

import (
	"github.com/kben19/CleoApp/pkg/domain"
	"github.com/kben19/CleoApp/pkg/lib/common"
	"github.com/kben19/CleoApp/pkg/types"
)

type UsecaseProduct struct {
	domainProduct  domain.DomainProductItf
	domainCategory domain.DomainCategoryItf
}

type UsecaseProductItf interface {
	GetProductByID(id int) (types.Product, error)
	GetProductListWithPagination(param types.StandardListParam) ([]types.Product, error)
}

func InitUsecaseProduct(product domain.DomainProductItf, category domain.DomainCategoryItf) UsecaseProduct {
	return UsecaseProduct{domainProduct: product, domainCategory: category}
}

func (u UsecaseProduct) GetProductByID(id int) (types.Product, error) {
	if id <= 0 {
		return types.Product{}, common.ErrInvalidIDParameter
	}

	product, err := u.domainProduct.GetProductByID(id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (u UsecaseProduct) GetProductListWithPagination(param types.StandardListParam) ([]types.Product, error) {
	products, err := u.domainProduct.GetProductListWithPagination(param)
	if err != nil {
		return products, err
	}

	categoryMap, err := u.domainCategory.GetCategoryMapDetail()
	if err != nil {
		return []types.Product{}, err
	}

	for i, val := range products {
		category := categoryMap[val.CategoryID]
		products[i].GroupName = category.GroupName
		products[i].TypeName = category.TypeName
	}
	return products, nil
}
