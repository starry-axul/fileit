package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/starry-axul/fileit/internal/client"
	"github.com/starry-axul/fileit/pkg/bootstrap"
	"github.com/starry-axul/fileit/pkg/handler"
	"log"
)

func main() {

	/*log := bootstrap.InitLogger()*/
	db, _, err := bootstrap.ConnectLocal()
	if err != nil {
		log.Fatal(err.Error())
	}

	r := client.NewRepository(db)
	s := client.NewService(r)
	e := client.MakeEndpoints(s, client.Config{})
	h := handler.NewLambdaClientRegen(e)
	lambda.StartHandler(h)
}
