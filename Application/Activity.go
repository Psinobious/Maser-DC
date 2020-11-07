package usecase

import (
	Domain "github.com/Psinobious/Maser-DC/Domain"
)

type ActivityInteractor struct {
	UserRepository       Domain.ClientRepository
	ConnectionRepository Domain.ConnectionRepository
	ActivityRepository   Domain.ActivityRepository
	ContentRepository    Domain.ContentRepository
	MaintainerRepository Domain.MaintainerRepository
	RoleRepository       Domain.RoleRepository
}

func (interactor *ActivityInteractor) Connections(activityID string) ([]Domain.Connection, error) {
	users, err := interactor.ConnectionRepository.FindUsersById(activityID)
	return users, err
}
func (interactor *ActivityInteractor) AddConnection(userId string, activityId string) error {
	user, err := interactor.UserRepository.FindById(userId)
	activity, err := interactor.ActivityRepository.FindById(activityId)
	connection := Domain.Connection{user.ClientID, activity.ActivityID, false}
	err = interactor.ConnectionRepository.Store(connection)
	return err
}
func (interactor *ActivityInteractor) RemoveConnection(userId string, activityId string) error {
	user, err := interactor.UserRepository.FindById(userId)
	activity, err := interactor.ActivityRepository.FindById(activityId)
	err = interactor.ConnectionRepository.Delete(user.ClientID, activity.ActivityID)
	return err
}
func (interactor *ActivityInteractor) FindConnection(userID string, activityID string) (Domain.Connection, error) {
	connection, err := interactor.ConnectionRepository.FindById(userID, activityID)
	return connection, err
}

func (interactor *ActivityInteractor) FindActivityById(ActivityID string) (Domain.Activity, error) {
	activity, err := interactor.ActivityRepository.FindById(ActivityID)
	maintainers, err := interactor.MaintainerRepository.FindUsers(ActivityID)
	activity.Maintainers = maintainers
	return activity, err
}
func (interactor *ActivityInteractor) AddActivity(Activity Domain.Activity, userID string, roleID string) error {
	maintainer := Domain.Maintainer{userID, Activity.ActivityID, roleID}
	err := interactor.ActivityRepository.Store(Activity)
	err = interactor.MaintainerRepository.Store(maintainer)
	return err
}
func (interactor *ActivityInteractor) RemoveActivity(ActivityID string) error {
	err := interactor.ConnectionRepository.Purge(ActivityID)
	err = interactor.ContentRepository.Purge(ActivityID)
	err = interactor.ActivityRepository.Delete(ActivityID)
	return err
}

func (interactor *ActivityInteractor) Maintainers(ActivityID string) ([]Domain.Maintainer, error) {
	users, err := interactor.MaintainerRepository.FindUsers(ActivityID)

	return users, err
}
func (interactor *ActivityInteractor) FindMaintainer(ActivityID string, UserID string) (Domain.Maintainer, error) {
	users, err := interactor.MaintainerRepository.FindUsers(ActivityID)
	user := Domain.Maintainer{}
	for i := 0; i < len(users); i++ {
		if users[i].UserID == UserID {
			user = users[i]
		}
	}
	return user, err
}
func (interactor *ActivityInteractor) ChangeRole(UserID string, activityID string, roleID string) error {
	maintainer, err := interactor.MaintainerRepository.FindById(UserID, activityID)
	role, err := interactor.RoleRepository.FindById(roleID)
	if err != nil {
		return err
	}
	maintainer.RoleID = role.RoleID
	err = interactor.MaintainerRepository.ChangeRole(maintainer)
	return err
}
func (interactor *ActivityInteractor) AddMaintainer(UserID string, ActivityID string, RoleID string) error {
	maintainer := Domain.Maintainer{UserID, ActivityID, RoleID}
	err := interactor.MaintainerRepository.Store(maintainer)
	return err
}
func (interactor *ActivityInteractor) RemoveMaintainer(userID string, activityID string) error {
	err := interactor.MaintainerRepository.Remove(userID, activityID)
	return err
}

func (interactor *ActivityInteractor) Contents(ActivityID string) ([]Domain.Content, error) {
	content, err := interactor.ContentRepository.FincByActivityID(ActivityID)
	return content, err
}
func (interactor *ActivityInteractor) FindContentById(ActivityID string, ContentID string) (Domain.Content, error) {
	content, err := interactor.ContentRepository.FindById(ContentID)
	return content, err
}
func (interactor *ActivityInteractor) AddContent(content Domain.Content) error {
	err := interactor.ContentRepository.Store(content)
	return err
}
func (interactor *ActivityInteractor) RemoveContent(ActivityID string, ContentID string) error {
	err := interactor.ContentRepository.Remove(ContentID, ActivityID)
	return err
}
