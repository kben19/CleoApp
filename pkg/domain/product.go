package domain

import (
	"database/sql"

	"github.com/kben19/CleoApp/pkg/types"
)

// ---------- DOMAIN ---------- //

type DomainProduct struct {
	rsc ResourceProductItf
}

type DomainProductItf interface {
	GetProductByID(id int) (types.Product, error)
}

func InitDomainProduct(db *sql.DB) DomainProduct {
	return DomainProduct{rsc: resourceProduct{db: db}}
}

func (d DomainProduct) GetProductByID(id int) (types.Product, error) {
	return d.rsc.GetProductByID(id)
}

// ---------- RESOURCE ---------- //

type resourceProduct struct {
	db *sql.DB
}

type ResourceProductItf interface {
	GetProductByID(id int) (types.Product, error)
}

func (r resourceProduct) GetProductByID(id int) (types.Product, error) {
	product := types.Product{}
	rows, err := r.db.Query(queryGetProductByID, id)
	if err != nil {
		return product, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Article, &product.Color, &product.Price, &product.BrandID)
		if err != nil {
			return product, err
		}
	}
	return product, nil
}
