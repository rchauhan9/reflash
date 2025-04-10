package auth

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func Middleware() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (_ interface{}, err error) {
			newContext := context.WithValue(ctx, "userID", "470af03f-886e-44e4-a42d-fb42140018e1")
			return next(newContext, request)
		}
	}
}
