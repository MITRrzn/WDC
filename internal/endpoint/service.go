package endpoint

import "WDC/internal/repository/endpoint"

type EndpointService struct {
	repo endpoint.Repository
}

func NewService(repo endpoint.Repository) *EndpointService {
	return &EndpointService{repo: repo}
}
