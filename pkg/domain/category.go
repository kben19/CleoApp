package domain

import (
	"database/sql"
	"github.com/kben19/CleoApp/pkg/types"
)

type DomainCategory struct {
	rsc ResourceCategoryItf
}

type DomainCategoryItf interface {
	GetCategoryMapDetail() (map[int]types.Category, error)
}

func InitDomainCategory(db *sql.DB) DomainCategory {
	return DomainCategory{rsc: resourceCategory{db: db}}
}

func (d DomainCategory) GetCategoryMapDetail() (map[int]types.Category, error) {
	productList, err := d.rsc.GetCategoryListDetail()
	if err != nil {
		return nil, err
	}
	productMap := make(map[int]types.Category)
	for i, val := range productList {
		productMap[val.ID] = productList[i]
	}
	return productMap, nil
}

type resourceCategory struct {
	db *sql.DB
}

type ResourceCategoryItf interface {
	GetCategoryListDetail() ([]types.Category, error)
}

func (r resourceCategory) GetCategoryListDetail() ([]types.Category, error) {
	rows, err := r.db.Query(queryGetCategoryListDetail)
	if err != nil {
		return nil, err
	}
	categories := []types.Category{}
	defer rows.Close()
	for rows.Next() {
		category := types.Category{}
		err := rows.Scan(
			&category.ID,
			&category.GroupID,
			&category.GroupName,
			&category.TypeID,
			&category.TypeName,
		)
		if err != nil {
			return []types.Category{}, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
