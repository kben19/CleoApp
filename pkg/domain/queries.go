package domain

var (
	// PRODUCT
	queryGetProductByID = `SELECT id, article, color, brand_id, category_id, image FROM product WHERE id = $1`

	queryGetProductPagination = `
		SELECT p.id, p.article, p.color, b.brand_name, p.category_id, p.image
		FROM product p
		INNER JOIN brand b ON b.id = p.brand_id
		ORDER BY %s %s
		OFFSET $1 LIMIT $2
		`

	// CATEGORY
	queryGetCategoryListDetail = `
		SELECT c.id, c.group_id, cg.group_name, c.type_id, ct.type_name
		FROM category c
		INNER JOIN category_group cg ON c.group_id = cg.id
		INNER JOIN category_type ct ON c.type_id = ct.id
		`
)
