package hello

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type sayHelloRequest struct {
	Name string `json:"name"`
}

type sayHelloResponse struct {
	Greeting string `json:"greeting"`
}

type Endpoints struct {
	SayHello endpoint.Endpoint
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		SayHello: MakeSayHelloEndpoint(svc),
	}
}

func MakeSayHelloEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(sayHelloRequest)
		greeting, err := svc.SayHello(ctx, req.Name)
		return sayHelloResponse{Greeting: greeting}, err
	}
}
