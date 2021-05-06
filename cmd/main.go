package main

import (
	bing_bang_boom "github.com/dontogbe/bing-bang-boom"
	"github.com/dontogbe/bing-bang-boom/bingbangboomsvc"
	kitHTTP "github.com/go-kit/kit/transport/http"
	"log"
	"net/http"
)

func main() {
	svc := bingbangboomsvc.NewSvc()
	serviceEndpoints := bing_bang_boom.NewServiceEndpoints(svc)
	mappingHandler := kitHTTP.NewServer(
		serviceEndpoints.GetDomainEndpoint,
		bing_bang_boom.DecodeMappingRequest,
		bing_bang_boom.EncodeMappingResponse,
	)
	http.Handle("/map", mappingHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
