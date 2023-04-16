package client

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/response"
	"github.com/ncostamagna/gocourse_meta/meta"
)

type (
	GetAllReq struct {
		Limit int    `json:"limit"`
		Page  int    `json:"page"`
		Token string `json:"authorization"`
	}

	CreateReq struct {
		Name  string `json:"name"`
		Token string `json:"authorization"`
	}

	CreateRes struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		ClientID   string `json:"client_id"`
		ReadToken  string `json:"read_token"`
		WriteToken string `json:"write_token"`
	}

	RegenReq struct {
		ID    string `json:"id"`
		Token string `json:"authorization"`
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
		Create Controller
		Regen  Controller
	}

	Config struct {
		LimPageDef string
	}
)

func MakeEndpoints(s Service, c Config) Endpoints {
	return Endpoints{
		GetAll: makeGetAllEndpoint(s, c),
		Create: makeCreateEndpoint(s),
		Regen:  makeRegenEndpoint(s),
	}
}

func makeGetAllEndpoint(s Service, c Config) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllReq)

		count, err := s.Count(ctx)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		meta, err := meta.New(req.Page, req.Limit, count, c.LimPageDef)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		cs, err := s.GetAll(ctx, meta.Offset(), meta.Limit())
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		return response.OK("success", cs, nil, nil), nil
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateReq)

		client, err := s.Create(ctx, req.Name)
		if err != nil {
			return nil, response.InternalServerError(err.Error())
		}

		res := CreateRes{
			ID:         client.ID,
			Name:       client.Name,
			ReadToken:  client.ReadToken,
			WriteToken: client.WriteToken,
		}
		return response.OK("success", res, nil, nil), nil
	}
}

func makeRegenEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return response.OK("success", request, nil, nil), nil
	}
}
