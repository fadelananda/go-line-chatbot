package entity

import "golang.org/x/oauth2"

type User struct {
	LineId    string        `dynamodbav:"line_id"`
	AuthToken *oauth2.Token `dynamodbav:"auth_token"`
	Email     string        `dynamodbav:"email"`
}

func (user User) IsEmpty() bool {
	return user.LineId == "" && user.Email == "" && user.AuthToken == nil
}
