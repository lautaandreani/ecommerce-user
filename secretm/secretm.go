package secretm

import (
	"encoding/json"
	"fmt"

	"home/lautaro/dev/ecommerce-user/awsgo"
	"home/lautaro/dev/ecommerce-user/models"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecrets(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println("> Getting secret" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	value, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*value.SecretString), &secretData)
	fmt.Println("> Successfully getting secret" + secretName)

	return secretData, nil
}
