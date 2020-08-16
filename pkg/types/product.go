package types

type Product struct {
	ID      int     `json:"id" db:"id"`
	Article string  `json:"article" db:"article"`
	Color   string  `json:"color" db:"color"`
	Price   float64 `json:"price" db:"price"`
	BrandID int     `json:"brand_id" db:"brand_id"`
}
