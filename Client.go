package domain

type Client struct {
	ClientID    string `json:"Username"`
	Email       string `json:"Email"`
	FirstName   string `json:"Firstname"`
	LastName    string `json:"Lastname"`
	Password    string `json:"Password"`
	Permissions []string
	Maintainers []Maintainer
	Connections []Connection
}
type ClientRepository interface {
	Store(client *Client) error
	FindById(ClientID string) (Client, error)
	CheckIfExist(email string) bool
	Delete(userID string) error
	Update(client Client) error
}
