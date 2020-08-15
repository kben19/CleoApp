package server

import "database/sql"

type usecases struct {
}

// Init function from usecases to resource
func Init(db *sql.DB) usecases {
	useCase := usecases{}

	return useCase
}
