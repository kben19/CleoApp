package types

type Product struct {
	ID         int    `json:"id" db:"id"`
	Article    string `json:"article" db:"article"`
	Color      string `json:"color" db:"color"`
	BrandID    int    `json:"brand_id" db:"brand_id"`
	CategoryID int    `json:"category_id" db:"brand_id"`
	Image      string `json:"image" db:"image"`
	BrandName  string `json:"brand_name" db:"brand_name"`
	BrandDesc  string `json:"brand_description" db:"brand_description"`
	GroupName  string `json:"group_name" db:"group_name"`
	TypeName   string `json:"type_name" db:"type_name"`
}

type Brand struct {
	ID        int    `json:"id" db:"id"`
	BrandName string `json:"brand_name" db:"brand_name"`
	BrandDesc string `json:"brand_description" db:"brand_description"`
}

type Category struct {
	ID        int    `json:"id" db:"id"`
	GroupID   int    `json:"group_id" db:"group_id"`
	TypeID    int    `json:"type_id" db:"type_id"`
	GroupName string `json:"group_name" db:"group_name"`
	TypeName  string `json:"type_name" db:"type_name"`
}
