package main

import (
	//"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/starry-axul/fileit/internal/client"
	//"github.com/starry-axul/fileit/pkg/bootstrap"
	"github.com/starry-axul/fileit/pkg/handler"
)

func main() {

	/*
	log := bootstrap.InitLogger()
	db, _, err := bootstrap.ConnectLocal(log)
	if err != nil {
		os.Exit(-1)
	}*/

	e := client.MakeEndpoints()
	h := handler.NewLambdaClientGetAll(e)
	lambda.StartHandler(h)
}
