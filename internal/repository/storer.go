package repository

import "github.com/fadelananda/go-line-chatbot/entity"

type UserStorer interface {
	ListUsers() ([]entity.User, error)
	AddUser(user entity.User) error
	GetUserById(id string) (entity.User, error)
}
