package handler

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalhouse-dev/dh-kit/request"
	"github.com/digitalhouse-dev/dh-kit/response"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/awslambda"
	"github.com/go-kit/log"
	"github.com/starry-axul/fileit/internal/client"
	"gorm.io/gorm"
)

func NewLambdaClientGetAll(endpoints client.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.GetAll), decodeGetAllHandler, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func NewLambdaClientCreate(endpoints client.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.Create), decodeCreateHandler, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func NewLambdaClientRegen(endpoints client.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.Regen), decodeRegenHandler, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func decodeGetAllHandler(_ context.Context, payload []byte) (interface{}, error) {
	var event events.APIGatewayProxyRequest
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, response.InternalServerError(err.Error())
	}

	var req client.GetAllReq
	if err := request.DecodeMap(event.Headers, &req); err != nil {
		return nil, response.BadRequest(err.Error())
	}
	err := request.DecodeMap(event.QueryStringParameters, &req)
	if err != nil {
		return nil, response.InternalServerError(err.Error())
	}
	return req, nil
}

func decodeCreateHandler(_ context.Context, payload []byte) (interface{}, error) {
	var event events.APIGatewayProxyRequest
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, response.InternalServerError(err.Error())
	}

	var req client.CreateReq
	if err := request.DecodeMap(event.QueryStringParameters, &req); err != nil {
		return nil, response.InternalServerError(err.Error())
	}
	if err := request.DecodeMap(event.Headers, &req); err != nil {
		return nil, response.BadRequest(err.Error())
	}
	if err := json.Unmarshal([]byte(event.Body), &req); err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return req, nil
}

func decodeRegenHandler(_ context.Context, payload []byte) (interface{}, error) {
	var event events.APIGatewayProxyRequest
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, response.InternalServerError(err.Error())
	}

	var req client.RegenReq
	if err := json.Unmarshal([]byte(event.Body), &req); err != nil {
		return nil, response.BadRequest(err.Error())
	}
	if err := request.DecodeMap(event.Headers, &req); err != nil {
		return nil, response.BadRequest(err.Error())
	}
	if err := request.DecodeMap(event.PathParameters, &req); err != nil {
		return nil, response.BadRequest(err.Error())
	}

	return req, nil
}

func EncodeResponse(_ context.Context, resp interface{}) ([]byte, error) {
	var res response.Response
	switch r := resp.(type) {
	case response.Response:
		res = r
	default:
		res = response.InternalServerError("unknown response type")
	}
	return APIGatewayProxyResponse(res)
}

func APIGatewayProxyResponse(res response.Response) ([]byte, error) {
	bytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	awsResponse := events.APIGatewayProxyResponse{
		Body:       string(bytes),
		StatusCode: res.StatusCode(),
		Headers:    res.GetHeaders(),
	}
	return json.Marshal(awsResponse)
}

func HandlerErrorEncoder(log log.Logger) awslambda.HandlerOption {
	return awslambda.HandlerErrorEncoder(
		awslambda.ErrorEncoder(errorEncoder(log)),
	)
}

func HandlerFinalizer(log log.Logger) func(context.Context, []byte, error) {
	return func(ctx context.Context, resp []byte, err error) {
		if err != nil {
			_ = log.Log("err", err)
		}
	}
}

func errorEncoder(log log.Logger) func(context.Context, error) ([]byte, error) {
	return func(_ context.Context, err error) ([]byte, error) {
		res := buildResponse(err, log)
		return APIGatewayProxyResponse(res)
	}
}

func buildResponse(err error, log log.Logger) response.Response {
	switch r := err.(type) {
	case response.Response:
		return r
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound("")
	}
	_ = log.Log("err", err)
	return response.InternalServerError("")
}
