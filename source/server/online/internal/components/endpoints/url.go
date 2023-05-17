package endpoints

import (
	"net/url"
	"online/internal/conf"
)

func NewEndPoints(cf *conf.Bootstrap) []*url.URL {
	var endpoints []*url.URL
	if cf.Server.Http.Endpoint != "" {
		httpEndpoint, _ := url.Parse("http://" + cf.Server.Http.Endpoint)
		endpoints = append(endpoints, httpEndpoint)
	}
	if cf.Server.Grpc.Endpoint != "" {
		grpcEndpoint, _ := url.Parse("grpc://" + cf.Server.Grpc.Endpoint)
		endpoints = append(endpoints, grpcEndpoint)
	}
	return endpoints
}
