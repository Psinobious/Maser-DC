package interfaces

import (
	Domain "github.com/Psinobious/Maser-DC/Domain"
)

type ActivityInteractor interface {
	Connections(ActivityID string) ([]Domain.Connection, error)
	AddConnection(userID string, activityID string) error
	RemoveConnection(userID string, activityID string) error
	FindConnection(userID string, activityID string) (Domain.Connection, error)

	FindActivityById(ActivityID string) (Domain.Activity, error)
	AddActivity(Activity Domain.Activity, userID string, roleID string) error
	RemoveActivity(ActivityID string, userID string) error
	UpdateActivity(Activity Domain.Activity, userID string) error

	Maintainers(ActivityID string) ([]Domain.Maintainer, error)
	FindMaintainer(ActivityID string, UserID string) (Domain.Maintainer, error)
	FindByRole(ActivityID string, Role string) (Domain.Maintainer, error)
	ChangeRole(UserID string, activityID string, roleID string) error
	AddMaintainer(Maintainer Domain.Maintainer) error
	RemoveMaintainer(userID string, activityID string) error

	Contents(ActivityID string) ([]Domain.Content, error)
	FindContentById(ActivityID string, ContentID string) (Domain.Content, error)
	FindContentByGroup(ActivityID string, Group string) ([]Domain.Content, error)
	AddContent(content Domain.Content, userID string) error
	RemoveContent(ActivityID string, ContentID string, userID string) error
}
