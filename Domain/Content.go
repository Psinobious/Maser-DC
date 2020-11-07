package domain

type Content struct {
	ContentID   string
	ActivityID  string
	Title       string
	Group       string
	DateCreated string
	Reference   string
}

type ContentRepository interface {
	FindById(contentID string) (Content, error)
	FincByActivityID(activityID string) ([]Content, error)
	Store(Content) error
	Purge(activityID string) error
	Remove(contentID string, activityID string) error
	Update(content Content) error
}
