package service

import (
	"database/sql"
)

// A TODOService implements CRUD of TODO entities.
type IndexService struct {
	db *sql.DB
}

// NewTODOService returns new TODOService.
func NewIndexService(db *sql.DB) *IndexService {
	return &IndexService{
		db: db,
	}
}
