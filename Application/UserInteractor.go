package interfaces

import (
	Domain "github.com/Psinobious/Maser-DC/Domain"
)

type UserInteractor interface {
	CreateUser(FirstName string, LastName string, Email string, Password string) error
	DeleteUser(UserID string) error
	FindUser(UserID string) (Domain.Client, error)
	ChangeFirstName(userID string, FirstName string) error
	ChangeLastName(userID string, LastName string) error
	ChangeEmail(userID string, email string) error
	ChangePassword(userID string, password string) error
}
