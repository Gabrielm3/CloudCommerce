package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielm3/cloudcommerce/awsgo"
)



func main(){
	lambda.Start(StartLambda)
}

func StartLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error){
	awsgo.Init()

	if !ValidateParameters() {
		fmt.Println("Missing parameters. Put SecretManager in environment variables.")
		err := errors.New("Missing parameters. Put SecretManager in environment variables.")
		return event, err
	}
}


func ValidateParameters() bool {
	var returnParam bool
	_, returnParam = os.LookupEnv("SecretName")
	return returnParam
}