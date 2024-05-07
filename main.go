package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/gabrielm3/cloudcommerce/awsgo"
	"github.com/gabrielm3/cloudcommerce/bd"
	"github.com/gabrielm3/cloudcommerce/models"
)



func main(){
	lambda.Start(StartLambda)
}

func StartLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error){
	awsgo.Init()

	if !ValidateParameters() {
		fmt.Println("Missing parameters. Put SecretManager in environment variables")
		err := errors.New("missing parameters. Put secretmanager in environment variables")
		return event, err
	}

	var dat models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
			case "email":
				dat.UserEmail = att
				fmt.Println("Email: " + dat.UserEmail)
			case "sub":
				dat.UserUUID = att
				fmt.Println("Sub: " + dat.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret: " + err.Error())
		return event, err
	}

	err = bd.SignUp(dat)
	return event, err
}


func ValidateParameters() bool {
	var returnParam bool
	_, returnParam = os.LookupEnv("SecretName")
	return returnParam
}