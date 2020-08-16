package domain

import (
	"database/sql"
	"fmt"

	"github.com/kben19/CleoApp/pkg/types"
)

// ---------- DOMAIN ---------- //

type DomainProduct struct {
	rsc ResourceProductItf
}

type DomainProductItf interface {
	GetProductByID(id int) (types.Product, error)
	GetProductListWithPagination(param types.StandardListParam) ([]types.Product, error)
}

func InitDomainProduct(db *sql.DB) DomainProduct {
	return DomainProduct{rsc: resourceProduct{db: db}}
}

func (d DomainProduct) GetProductByID(id int) (types.Product, error) {
	return d.rsc.GetProductByID(id)
}

func (d DomainProduct) GetProductListWithPagination(param types.StandardListParam) ([]types.Product, error) {
	return d.rsc.GetProductListWithPagination(param)
}

// ---------- RESOURCE ---------- //

type resourceProduct struct {
	db *sql.DB
}

type ResourceProductItf interface {
	GetProductByID(id int) (types.Product, error)
	GetProductListWithPagination(param types.StandardListParam) ([]types.Product, error)
}

func (r resourceProduct) GetProductByID(id int) (types.Product, error) {
	product := types.Product{}
	rows, err := r.db.Query(queryGetProductByID, id)
	if err != nil {
		return product, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&product.ID,
			&product.Article,
			&product.Color,
			&product.BrandID,
			&product.CategoryID,
			&product.Image,
		)
		if err != nil {
			return product, err
		}
	}
	return product, nil
}

func (r resourceProduct) GetProductListWithPagination(param types.StandardListParam) ([]types.Product, error) {
	products := []types.Product{}
	query := fmt.Sprintf(queryGetProductPagination, param.OrderBy, param.OrderType)
	rows, err := r.db.Query(query, param.Offset, param.Limit)
	if err != nil {
		return products, err
	}
	defer rows.Close()
	for rows.Next() {
		product := types.Product{}
		err := rows.Scan(
			&product.ID,
			&product.Article,
			&product.Color,
			&product.BrandName,
			&product.CategoryID,
			&product.Image,
		)
		if err != nil {
			return []types.Product{}, err
		}
		products = append(products, product)
	}
	return products, nil
}
