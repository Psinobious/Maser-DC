package domain

type Maintainer struct {
	UserID     string
	ActivityID string
	RoleID     string
}
type MaintainerRepository interface {
	FindUsers(ActivityID string) ([]Maintainer, error)
	FindActivities(UserID string) ([]Maintainer, error)
	FindById(UserID string, ActivityID string) (Maintainer, error)
	Store(Maintainer) error
	Remove(UserID string, ActivityID string) error
	Purge(UserID string) error
	ChangeRole(maintainer Maintainer) error
}
