package domain

type Connection struct {
	ClientID      string `json:"ClientID"`
	ActivityID    string `json:"ActivityID"`
	Notifications bool   `json:"Notification"`
}
type ConnectionRepository interface {
	Store(connection Connection) error
	Delete(ClientID string, ActivityID string) error
	Purge(ClientID string) error
	Update(connection Connection) error
	FindUsersById(ActivityID string) ([]Connection, error)
	FindActivitiesById(UserID string) ([]Connection, error)
	FindById(ClientID string, ActivityID string) (Connection, error)
}
