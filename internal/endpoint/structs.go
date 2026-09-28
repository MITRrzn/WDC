package endpoint

import "time"

type Input struct {
	Url    string `json:"url"`
	Secret string `json:"secret"`
}

type Response struct {
	Status string         `json:"status"`
	Data   EndpointStruct `json:"data"`
}

type EndpointStruct struct {
	Url       string
	Secret    string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
