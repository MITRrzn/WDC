package endpoint

import "context"

type EndpointRepository interface {
	Create(ctx context.Context, endpoint EndpointStruct) error
}

type ServiceInterface interface {
	Create(ctx context.Context) error
}
