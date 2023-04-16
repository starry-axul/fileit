package client

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	GetAllReq struct {
		Token string `json:"authorization"`
	}

	CreateReq struct {
		Name string `json:"name"`
		Token string `json:"authorization"`
	}

	CreateRes struct {
		ID string `json:"id"`
		Name string `json:"name"`
		ClientID string `json:"client_id"`
		ReadToken string `json:"read_token"`
		WriteToken string `json:"write_token"`
	}

	RegenReq struct {
		ID string `json:"id"`
		Token string `json:"authorization"`
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
		Create Controller
		Regen  Controller
	}
)

func MakeEndpoints() Endpoints {
	return Endpoints{
		GetAll: makeGetAllEndpoint(),
		Create: makeCreateEndpoint(),
		Regen: makeRegenEndpoint(),

	}
}

func makeGetAllEndpoint() Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return response.OK("success", request, nil, nil), nil
	}
}

func makeCreateEndpoint() Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return response.OK("success", request, nil, nil), nil
	}
}

func makeRegenEndpoint() Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return response.OK("success", request, nil, nil), nil
	}
}