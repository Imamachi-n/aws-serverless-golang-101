package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type TestResponse struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

// Handler function Using AWS Lambda Proxy Request
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Get the path parameter that was sent
	name := request.PathParameters["name"]

	//Get the query parameter that was sent
	msg := request.QueryStringParameters["msg"]

	test := TestResponse{
		Name: name,
		Msg:  msg,
	}
	jsonRes, err := json.Marshal(test)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	}

	//Generate message that want to be sent as body
	// message := fmt.Sprintf(" { \"Message\" : \"Hello %s %s \" } ", name, msg)

	//Returning response with AWS Lambda Proxy Response
	return events.APIGatewayProxyResponse{
		Body:       string(jsonRes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
