package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielm3/cloudcommerce/awsgo"
	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/handlers"
)

func main(){
	lambda.Start(StartLambda)
}

func StartLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error){
	awsgo.Init()

	if !ValidateParameters() {
		panic("Error: Missing parameters")
	}

	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body: string(message),
		Headers: headersResp,
	}

	return res, nil
}


func ValidateParameters() bool {
	_, returnParam := os.LookupEnv("SecretName")
	if !returnParam {
		return returnParam
	}

	_, returnParam = os.LookupEnv("UrlPrefix")
	if !returnParam {
		return returnParam
	}

	return returnParam
}