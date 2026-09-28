package endpoint

import "context"

type Endpoint interface {
	Create(ctx context.Context, endpoint Endpoint) error
}
