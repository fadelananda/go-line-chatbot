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
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/fadelananda/go-line-chatbot/entity"
)

type AWSClient struct {
	dynamoDbClient *dynamodb.Client
}

type awsClientError struct {
	FunctionName string
	Err          error
}

func (e *awsClientError) Error() string {
	return fmt.Sprintf("Google client error from %s, error: %v", e.FunctionName, e.Err)
}

func newAWSClientError(functionName string, err error) *awsClientError {
	return &awsClientError{
		FunctionName: functionName,
		Err:          err,
	}
}

func NewAWSClient() (*AWSClient, error) {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if accessKey == "" || secretKey == "" {
		return nil, newAWSClientError("NewAWSClient", errors.New("AWS access key or secret key not provided"))
	}
	creds := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion("ap-southeast-1"),
	)
	if err != nil {
		return nil, newAWSClientError("NewAWSClient", err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	return &AWSClient{
		dynamoDbClient: svc,
	}, nil
}

func (client *AWSClient) AddUser(user entity.User) error {
	item, err := attributevalue.MarshalMap(user)
	if err != nil {
		return newAWSClientError("AddUser", err)
	}

	_, err = client.dynamoDbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String("user"), Item: item,
	})
	if err != nil {
		return newAWSClientError("AddUser", err)
	}
	return nil
}

func (client *AWSClient) GetDataByLineId(lineId string) (entity.User, error) {
	user := entity.User{LineId: lineId}

	key := map[string]types.AttributeValue{
		"line_id": &types.AttributeValueMemberS{Value: lineId},
	}
	response, err := client.dynamoDbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		Key: key, TableName: aws.String("user"),
	})
	if err != nil {
		return entity.User{}, newAWSClientError("GetDataByLineId", err)
	}
	if len(response.Item) == 0 {
		return entity.User{}, nil
	}

	err = attributevalue.UnmarshalMap(response.Item, &user)
	if err != nil {
		return entity.User{}, newAWSClientError("GetDataByLineId", err)
	}

	return user, err
}

// TODO: error handling since this is not tested
func (client *AWSClient) UpdateUser(lineId string, updateData entity.User) (map[string]map[string]interface{}, error) {
	var attributeMap map[string]map[string]interface{}

	key := map[string]types.AttributeValue{
		"line_id": &types.AttributeValueMemberS{Value: lineId},
	}

	updateExpression := expression.Set(expression.Name("email"), expression.Value(updateData.Email))
	expr, err := expression.NewBuilder().WithUpdate(updateExpression).Build()
	if err != nil {
		log.Printf("Couldn't build expression for update. Here's why: %v\n", err)
	} else {
		response, err := client.dynamoDbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
			TableName:                 aws.String("user"),
			Key:                       key,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			UpdateExpression:          expr.Update(),
			ReturnValues:              types.ReturnValueUpdatedNew,
		})
		if err != nil {
			log.Printf("Couldn't update movie %v. Here's why: %v\n", updateData.LineId, err)
		} else {
			err = attributevalue.UnmarshalMap(response.Attributes, &attributeMap)
			if err != nil {
				log.Printf("Couldn't unmarshall update response. Here's why: %v\n", err)
			}
		}
	}
	fmt.Println(attributeMap)
	return attributeMap, err
}
