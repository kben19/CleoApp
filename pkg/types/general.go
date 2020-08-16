package types

// StandardListParam struct for general list param
type StandardListParam struct {
	Offset    int    `json:"offset" db:"offset"`
	Limit     int    `json:"limit" db:"limit"`
	OrderBy   string `json:"order_by"`
	OrderType string `json:"order_type"`
}
