package endpoint

import "context"

type EndpointRepository interface {
	Create(ctx context.Context, input Input) (EndpointStruct, error)
	Get(ctx context.Context, id int64) (EndpointStruct, error)
}

type ServiceInterface interface {
	Create(ctx context.Context, input Input) (EndpointStruct, error)
	Get(ctx context.Context, id int64) (EndpointStruct, error)
}
