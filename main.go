package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(pushMessage)
}

func pushMessage(request events.APIGatewayProxyRequest) (Chat, error) {
	name := "カッキー" // request.PathParameters["user_name"]
	var message string

	if request.PathParameters["text"] != "" {
		message = request.PathParameters["text"]
	} else {
		message = "わーい"
	}

	chat := Chat{
		Username: name,
		Text:     message,
	}

	return chat, nil
}

// Chat is a message.
type Chat struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}
