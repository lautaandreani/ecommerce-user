package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"home/lautaro/dev/ecommerce-user/awsgo"
	"home/lautaro/dev/ecommerce-user/db"
	"home/lautaro/dev/ecommerce-user/models"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(initLambda)
}

func initLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitAWS()

	if !existSecretName() {
		fmt.Println("Error when try access to 'SecretName'")
		err := errors.New("Error when try access to 'SecretName' please send value")
		return event, err
	}

	var data models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error reading secret" + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err
}

func existSecretName() bool {
	var getParameters bool
	_, getParameters = os.LookupEnv("SecretName")
	return getParameters
}
