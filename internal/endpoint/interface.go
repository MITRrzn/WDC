package endpoint

import "context"

type EndpointRepository interface {
	Create(ctx context.Context, endpoint Endpoint) error
}
