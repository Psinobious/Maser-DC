package domain

type Activity struct {
	ActivityID   string `json:"ActivityID"`
	OwnerID      string `json:"OwnerID"`
	Title        string `json:"Title"`
	ActivityType string `json:"ActivityType"`
	DateCreated  string `json:"Date"`
	Maintainers  []Maintainer
}
type ActivityRepository interface {
	Store(activity Activity) error
	Delete(ActivityID string) error
	FindById(activityID string) (Activity, error)
	Update(activity Activity) error
}
