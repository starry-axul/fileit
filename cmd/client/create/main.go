package main

import (
	//"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/starry-axul/fileit/internal/user"
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

	e := user.MakeEndpoints()
	h := handler.NewLambdaUserGetAll(e)
	lambda.StartHandler(h)
}
