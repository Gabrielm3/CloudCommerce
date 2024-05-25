package secretmg

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/gabrielm3/cloudcommerce/awsgo"
	"github.com/gabrielm3/cloudcommerce/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	fmt.Println("Getting secret: " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	result, err := svc.GetSecretValue(awsgo.Ctx, input)
	if err != nil {
		fmt.Println("Error retrieving secret: " + err.Error())
	}

	var secret models.SecretRDSJson
	json.Unmarshal([]byte(*result.SecretString), &secret)

	fmt.Println("Secret retrieved: " + secret.Username)

	return secret, nil
}