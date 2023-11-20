package models

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
)

type Environment struct {
	CredFileName   string
	ConfigFileName string
	Profile        string
	Region         string
}

func SetDefaultEnv() Environment {
	env := Environment{
		CredFileName:   "/root/.aws/credentials",
		ConfigFileName: "/root/.aws/config",
		Profile:        "default",
		Region:         "us-east-1",
	}
	return env
}

func GetAWSConfig() aws.Config {
	environ := SetDefaultEnv()

	conf, _ := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithRegion(environ.Region),
		awsConfig.WithSharedConfigProfile(environ.Profile),
		awsConfig.WithSharedCredentialsFiles([]string{
			environ.CredFileName,
		}),
	)

	return conf
}
