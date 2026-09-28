package endpoint

import (
	"WDC/internal/helper"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type EndpointHandler struct {
	service ServiceInterface
}

func NewHandler(service ServiceInterface) *EndpointHandler {
	return &EndpointHandler{
		service: service,
	}
}

func (h *EndpointHandler) CreateEndpoint(w http.ResponseWriter, r *http.Request) {
	var input Input
	decodeErr := json.NewDecoder(r.Body).Decode(&input)

	if decodeErr != nil {
		log.Println("Create endpoint, error decode input data", decodeErr)
		http.Error(w, decodeErr.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.service.Create(r.Context(), input)
	if err != nil {
		var validationErr ValidationError
		if errors.As(err, &validationErr) {
			helper.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		helper.WriteErrorResponse(w, "something goes wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encodeErr := json.NewEncoder(w).Encode(Response{
		Status: "OK",
		Data: ResponseData{
			Url:       result.Url,
			IsActive:  result.IsActive,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		},
	})
	if encodeErr != nil {
		log.Println("encode response error:", encodeErr)
	}
}
