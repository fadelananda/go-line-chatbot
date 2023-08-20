package client

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"golang.org/x/oauth2"
)

type AWSClient struct {
	dynamoDbClient *dynamodb.Client
}

type User struct {
	LineId    string        `dynamodbav:"line_id"`
	AuthToken *oauth2.Token `dynamodbav:"auth_token"`
	Email     string        `dynamodbav:"email"`
}

func NewAWSClient() (*AWSClient, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if accessKey == "" || secretKey == "" {
		return nil, errors.New("AWS access key or secret key not provided")
	}
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	// Create a config with credentials and region
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion("ap-southeast-1"), // Specify your desired region
	)
	if err != nil {
		fmt.Println("Error creating AWS config:", err)
		return nil, err
	}

	// Create a DynamoDB client
	svc := dynamodb.NewFromConfig(cfg)
	fmt.Println("========")
	fmt.Println(svc)
	fmt.Println("========")

	return &AWSClient{
		dynamoDbClient: svc,
	}, nil
}

func (client *AWSClient) ListTables() ([]string, error) {
	var tableNames []string
	fmt.Println("********")
	fmt.Println(client)
	fmt.Println("********")
	fmt.Println(client.dynamoDbClient)
	fmt.Println("--------")
	tables, err := client.dynamoDbClient.ListTables(
		context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		log.Printf("Couldn't list tables. Here's why: %v\n", err)
	} else {
		tableNames = tables.TableNames
	}
	fmt.Println(tableNames)
	return tableNames, err
}

func (client *AWSClient) AddUser(user User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		fmt.Println("123123")
	}

	_, err = client.dynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("user"), Item: item,
	})
	return err
}

func (client *AWSClient) GetDataByLineId(lineId string) (User, error) {
	user := User{LineId: lineId}
	fmt.Println(user)
	key := map[string]types.AttributeValue{
		"line_id": &types.AttributeValueMemberS{Value: "Ucbbee99b74cb44198c06181ce84a0b08"},
	}
	response, err := client.dynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: key, TableName: aws.String("user"),
	})
	if err != nil {
		log.Printf("Couldn't get info about %v. Here's why: %v\n", lineId, err)
	} else {
		err = attributevalue.UnmarshalMap(response.Item, &user)
		if err != nil {
			log.Printf("Couldn't unmarshal response. Here's why: %v\n", err)
		}
	}
	return user, err
}
