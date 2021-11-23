package bing_bang_boom

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"log"
	"net/http"

	"github.com/dontogbe/bing-bang-boom/bingbangboomsvc"
)

type serviceEndpoints struct {
	GetDomainEndpoint endpoint.Endpoint
}

func NewServiceEndpoints(svc bingbangboomsvc.Svc) *serviceEndpoints {
	return &serviceEndpoints{
		GetDomainEndpoint: makeGenerateMappingEndpoint(svc),
	}
}

func makeGenerateMappingEndpoint(svc bingbangboomsvc.Svc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(mappingRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		domain, err := svc.GenerateMapping(req.ID)
		if err != nil {
			return mappingResponse{Domain: domain}, err
		}
		return mappingResponse{Domain: domain}, nil
	}
}

type mappingRequest struct {
	ID int64 `json:"id"`
}

type mappingResponse struct {
	Domain []bingbangboomsvc.Domain `json:"domain"`
}

func DecodeMappingRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request mappingRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println(err)
		iErr := errors.New("Please provide numeric ID")
		return nil, iErr
	}
	return request, nil
}
func EncodeMappingResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
