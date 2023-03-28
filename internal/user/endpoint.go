package user

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	GetAllReq struct {
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		GetAll: makeGetAllEndpoint(),
	}
}

func makeGetAllEndpoint() Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return response.OK("success", request, nil, nil), nil
	}
}