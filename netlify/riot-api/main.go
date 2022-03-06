package main

import (
	"github.com/artlovecode/wordlists.tech/pkg/handlers"
	"github.com/artlovecode/wordlists.tech/pkg/service"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
)

type environment struct {
	POSTGRES_DB             string `envconfig:"POSTGRES_DB" required:"true"`
	POSTGRES_DB_USER        string `envconfig:"POSTGRES_DB_USER" required:"true"`
	POSTGRES_DB_PASSWORD    string `envconfig:"POSTGRES_DB_PASSWORD" required:"true"`
	RIOT_DATADRAGON_VERSION string `envconfig:"RIOT_DATADRAGON_URL" default:"12.5.1"`
}

func main() {
	var env environment
	envconfig.MustProcess("", &env)

	service := service.New(env.POSTGRES_DB, env.RIOT_DATADRAGON_VERSION)
	champListHandler := handlers.NewChampionListHandler(service)

	lambda.Start(champListHandler)
}
