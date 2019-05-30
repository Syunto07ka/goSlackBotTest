package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(pushMessage)
}

func pushMessage(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := "カッキー"   // request.PathParameters["name"]
	message := "わーい" // request.QueryStringParameters["message"]

	chat := Chat{
		Username: name,
		Text:     message,
	}

	jsonMessage, err := json.Marshal(chat)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonMessage),
		StatusCode: 200,
	}, nil
}

// Chat is a message.
type Chat struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}
