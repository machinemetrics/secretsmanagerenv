package aws

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"os"
)

func GetSecretData(name, region string) (map[string]interface{}, error) {
	var secrets map[string]interface{}
	// Grab env vars from Secrets Manager
	session := getSession(region)
	svc := secretsmanager.New(session)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(*result.SecretString), &secrets)
	if err != nil {
		return nil, err
	}

	return secrets, nil
}

func getSession(region string) *session.Session {
	if len(os.Getenv("AWS_SDK_LOAD_CONFIG")) == 0 {
		os.Setenv("AWS_SDK_LOAD_CONFIG", "TRUE")
		defer os.Unsetenv("AWS_SDK_LOAD_CONFIG")
	}

	config := aws.Config{}
	if len(region) > 0 {
		config.Region = aws.String(region)
	}
	return session.Must(session.NewSession(&config))
}
