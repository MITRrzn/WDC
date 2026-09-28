package endpoint

import (
	"context"
	"net/url"
)

type EndpointService struct {
	repo EndpointRepository
}

func NewService(repo EndpointRepository) *EndpointService {
	return &EndpointService{
		repo: repo,
	}
}

func (e EndpointService) Create(ctx context.Context, input Input) (EndpointStruct, error) {
	err := validateInput(input)
	if err != nil {
		return EndpointStruct{}, err
	}

	result, err := e.repo.Create(ctx, input)
	if err != nil {
		return EndpointStruct{}, err
	}

	return result, nil
}

func validateInput(input Input) error {
	if input.Url == "" {
		return ValidationError{Message: "url is required"}
	}

	u, err := url.Parse(input.Url)
	if err != nil {
		return ValidationError{Message: "url is invalid"}
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return ValidationError{Message: "url scheme must be http or https"}
	}
	if u.Host == "" {
		return ValidationError{Message: "url host is required"}
	}

	if input.Secret == "" {
		return ValidationError{Message: "secret is required"}
	}

	return nil
}
