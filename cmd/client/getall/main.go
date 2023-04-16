package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/starry-axul/fileit/internal/client"
	"github.com/starry-axul/fileit/pkg/bootstrap"
	"github.com/starry-axul/fileit/pkg/handler"
	"log"
	"os"
)

func main() {

	/*log := bootstrap.InitLogger()*/
	db, _, err := bootstrap.ConnectLocal()
	if err != nil {
		log.Fatal(err.Error())
	}

	pagLimDef := os.Getenv("PAGINATOR_LIMIT_DEFAULT")
	if pagLimDef == "" {
		log.Fatal("paginator limit default is required")
	}
	r := client.NewRepository(db)
	s := client.NewService(r)
	e := client.MakeEndpoints(s, client.Config{LimPageDef: pagLimDef})
	h := handler.NewLambdaClientGetAll(e)
	lambda.StartHandler(h)
}
