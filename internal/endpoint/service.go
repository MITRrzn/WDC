package endpoint

import "context"

type EndpointService struct {
	repo EndpointRepository
}

func NewService(repo EndpointRepository) *EndpointService {
	return &EndpointService{
		repo: repo,
	}
}

func (e EndpointService) Create(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
