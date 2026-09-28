package endpoint

import "net/http"

type EndpointHandler struct {
	service EndpointService
}

func NewHandler(service EndpointService) *EndpointHandler {
	return &EndpointHandler{
		service: service,
	}
}

func (h *EndpointHandler) CreateEndpoint(w http.ResponseWriter, r *http.Request) {}
