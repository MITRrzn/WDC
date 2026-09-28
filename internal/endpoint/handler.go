package endpoint

import (
	"WDC/internal/helper"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
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
		http.Error(w, "invalid data", http.StatusBadRequest)
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

func (h *EndpointHandler) Get(w http.ResponseWriter, r *http.Request) {
	reqId := r.PathValue("id")
	id, parseErr := strconv.ParseInt(reqId, 10, 64)
	if parseErr != nil {
		log.Println("invalid endpoint id:", parseErr)
		helper.WriteErrorResponse(w, "invalid endpoint id", http.StatusBadRequest)
		return
	}

	result, err := h.service.Get(r.Context(), id)

	if err != nil {
		var validationErr ValidationError
		if errors.As(err, &validationErr) {
			helper.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("get response error:", err)
		helper.WriteErrorResponse(w, "something goes wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
