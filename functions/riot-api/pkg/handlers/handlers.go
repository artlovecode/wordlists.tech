package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/artlovecode/wordlists.tech/functions/riot-api/pkg/service"
	"github.com/aws/aws-lambda-go/events"
)

func NewChampionListHandler(s service.Service) func(context.Context, events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return func(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
		data, err := s.GetData()
		if err != nil {
			fmt.Printf("Got error fetching champion data: %v", err)
			return &events.APIGatewayProxyResponse{
				StatusCode: http.StatusInternalServerError,
				Body:       "internal server error",
			}, nil
		}
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       string(data),
		}, nil
	}
}
