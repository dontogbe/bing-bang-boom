package bing_bang_boom

import (
	"context"
	"errors"
	"github.com/dontogbe/bing-bang-boom/bingbangboomsvc"
	"github.com/go-kit/kit/endpoint"
)

type serviceEndpoints struct {
	GetDomainEndpoint endpoint.Endpoint
}

func  NewServiceEndpoints(svc bingbangboomsvc.Svc) *serviceEndpoints {
	return &serviceEndpoints{
		GetDomainEndpoint: makeGenerateMappingEndpoint(svc),
	}
}

func makeGenerateMappingEndpoint(svc bingbangboomsvc.Svc) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req,ok:=request.(mappingRequest)
		if !ok{
			return nil,errors.New("invalid request")
		}
		domain,err:=svc.GenerateMapping(req.id)
		if err!=nil{
			return mappingResponse{Domain: domain},err
		}
		return mappingResponse{Domain: domain},nil
	}
}

type mappingRequest struct{
	id int64 `json:"id"`
}

type mappingResponse struct {
	Domain []bingbangboomsvc.Domain	`json:"domain"`
}