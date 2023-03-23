package outerServices

import (
	"auth/internal/entities"
	"context"
	"encoding/json"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"log"
	"os"
)

func GetSecret() (entities.AWSCredentials, error) {
	var c entities.AWSCredentials

	secretName := os.Getenv("SECRET_NAME")
	region := os.Getenv("REGION")

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			os.Getenv("ACCESS_KEY_ID"),
			os.Getenv("SECRET_ACCESS_KEY"),
			"",
		)),
	)
	if err != nil {
		log.Fatal(err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(os.Getenv("VERSION_STAGE")),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}

	var secretString = *result.SecretString

	err = json.Unmarshal([]byte(secretString), &c)

	if err != nil {
		return entities.AWSCredentials{}, err
	}

	return c, nil
}
