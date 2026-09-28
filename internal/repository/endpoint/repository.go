package endpoint

import (
	"context"
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

type EndpointStruct struct {
	name string
}

func (r Repository) Create(ctx context.Context, endpoint EndpointStruct) error {
	//TODO implement me
	panic("implement me")
}
