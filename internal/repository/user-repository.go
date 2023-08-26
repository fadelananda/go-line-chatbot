package repository

import (
	"github.com/fadelananda/go-line-chatbot/entity"
	"github.com/fadelananda/go-line-chatbot/internal/client"
)

type UserRepository struct {
	awsClient *client.AWSClient
}

func NewUserRepository(awsClient *client.AWSClient) *UserRepository {
	return &UserRepository{
		awsClient: awsClient,
	}
}

// TODO: implement
func (r *UserRepository) ListUsers() ([]entity.User, error) {
	return nil, nil
}

func (r *UserRepository) AddUser(user entity.User) error {
	return r.awsClient.AddUser(user)
}

func (r *UserRepository) GetUserById(id string) (entity.User, error) {
	return r.awsClient.GetDataByLineId(id)
}
