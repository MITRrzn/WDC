package endpoint

type EndpointService struct {
	repo EndpointRepository
}

func NewService(repo EndpointRepository) *EndpointService {
	return &EndpointService{
		repo: repo,
	}
}
