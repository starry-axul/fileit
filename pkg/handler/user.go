package handler

import (
	"context"
	"encoding/json"
	"errors"
	
	"github.com/aws/aws-lambda-go/events"
	"github.com/digitalhouse-dev/dh-kit/request"
	"github.com/digitalhouse-dev/dh-kit/response"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport/awslambda"
	"github.com/starry-axul/fileit/internal/user"
	"gorm.io/gorm"
)

func NewLambdaUserGetAll(endpoints user.Endpoints) *awslambda.Handler {
	return awslambda.NewHandler(endpoint.Endpoint(endpoints.GetAll), decodeGetAllHandler, EncodeResponse,
		HandlerErrorEncoder(nil), awslambda.HandlerFinalizer(HandlerFinalizer(nil)))
}

func decodeGetAllHandler(_ context.Context, payload []byte) (interface{}, error) {
	var event events.APIGatewayProxyRequest
	if err := json.Unmarshal(payload, &event); err != nil {
		return nil, response.InternalServerError(err.Error())
	}

	var req user.GetAllReq

	err := request.DecodeMap(event.QueryStringParameters, &req)
	if err != nil {
		return nil, response.InternalServerError(err.Error())
	}
	return req, nil
}

func EncodeResponse(_ context.Context, resp interface{}) ([]byte, error) {
	var res response.Response
	switch resp.(type) {
	case response.Response:
		res = resp.(response.Response)
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
			log.Log("err", err)
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
	switch err.(type) {
	case response.Response:
		return err.(response.Response)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.NotFound("")
	}
	log.Log("err", err)
	return response.InternalServerError("")
}
