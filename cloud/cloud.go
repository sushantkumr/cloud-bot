package cloud

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var ssmClient *ssm.SSM

func initializeAwsSession() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"),
	})

	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}

	ssmClient = ssm.New(sess)
}

func getParamtersFromSSM(parameterName string) string {
	initializeAwsSession()

	input := &ssm.GetParameterInput{
		Name:           aws.String(parameterName),
		WithDecryption: aws.Bool(true), // Set to true if the parameter is encrypted.
	}

	result, err := ssmClient.GetParameter(input)
	if err != nil {
		fmt.Println("Error getting parameter:", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fetched AWS Parameter: Name: %s\n", *result.Parameter.Name)
	return *result.Parameter.Value
}
