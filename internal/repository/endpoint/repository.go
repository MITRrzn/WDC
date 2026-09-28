package endpoint

import (
	"WDC/internal/endpoint"
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ctx context.Context, endpoint endpoint.EndpointStruct) error {
	//TODO implement me
	panic("implement me")
}
