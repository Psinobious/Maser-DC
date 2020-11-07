package domain

type Item struct {
	ItemID       string
	ActivityID   string
	OwnerID      string
	DateCreated  string
	DateModified string
	Group        string
	Reference    string
	Description  string
}
type ItemRepository interface {
	findItemByUser(OwnerID string) ([]Item, error)
	update(item Item) error
	remove(ItemID string) error
	purge(OwnerID string) error
}
