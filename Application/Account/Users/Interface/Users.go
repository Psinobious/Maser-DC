package usecase

import (
	Domain "github.com/Psinobious/Maser-DC/Domain"
)

type UserInteractor struct {
	UserRepository        Domain.ClientRepository
	MaintainerRepository  Domain.MaintainerRepository
	PermissionsRepository Domain.PermissionsRepository
	RoleRepository        Domain.RoleRepository
	ConnectionRepository  Domain.ConnectionRepository
}

func (interactor *UserInteractor) FindUser(UserID string) (Domain.Client, error) {
	user, err := interactor.UserRepository.FindById(UserID)
	permissions, err := interactor.PermissionsRepository.Permissions(UserID)
	maintainers, err := interactor.MaintainerRepository.FindActivities(UserID)
	connections, err := interactor.ConnectionRepository.FindActivitiesById(UserID)
	user.Permissions = permissions
	user.Maintainers = maintainers
	user.Connections = connections
	return user, err
}
func (interactor *UserInteractor) CreateUser(ClientID string, FirstName string, LastName string, Email string, Password string) error {
	var user Domain.Client

	user.ClientID = ClientID
	user.FirstName = FirstName
	user.LastName = LastName
	user.Email = Email
	user.Password = Password
	err := interactor.UserRepository.Store(&user)

	return err
}
func (interactor *UserInteractor) DeleteUser(userID string) error {
	err := interactor.MaintainerRepository.Purge(userID)
	err = interactor.ConnectionRepository.Purge(userID)
	err = interactor.UserRepository.Delete(userID)
	return err
}
func (interactor *UserInteractor) UpdateClient(User *Domain.Client) error {
	user, err := interactor.UserRepository.FindById(User.ClientID)
	err = interactor.UserRepository.Update(user)
	return err
}
func (interactor *UserInteractor) ChangePassword(userID string, password string) error {
	user, err := interactor.UserRepository.FindById(userID)
	user.Password = password

	err = interactor.UserRepository.Update(user)

	return err
}
