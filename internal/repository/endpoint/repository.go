package endpoint

import (
	"WDC/internal/endpoint"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ctx context.Context, input endpoint.Input) (endpoint.EndpointStruct, error) {
	var result endpoint.EndpointStruct

	err := r.db.QueryRowContext(
		ctx,
		`INSERT INTO endpoints (url, secret, is_active) VALUES ($1, $2, $3) RETURNING url, secret, is_active, created_at, updated_at`,
		input.Url,
		input.Secret,
		true,
	).Scan(&result.Url, &result.Secret, &result.IsActive, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		return endpoint.EndpointStruct{}, fmt.Errorf("create endpoint error: %w", err)
	}

	return result, nil
}

func (r Repository) Get(ctx context.Context, id int64) (endpoint.EndpointStruct, error) {
	var result endpoint.EndpointStruct

	err := r.db.QueryRowContext(
		ctx,
		`SELECT url, is_active, created_at, updated_at FROM endpoints WHERE id = $1`,
		id,
	).Scan(&result.Url, &result.IsActive, &result.CreatedAt, &result.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return endpoint.EndpointStruct{}, endpoint.NotFoundError{Message: "endpoint not found"}
		}
		return endpoint.EndpointStruct{}, fmt.Errorf("get endpoint error: %w", err)
	}

	return result, nil
}
