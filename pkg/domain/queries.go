package domain

var (
	queryGetProductByID = `SELECT id, article, color, price, brand_id FROM product WHERE id = $1`
)
