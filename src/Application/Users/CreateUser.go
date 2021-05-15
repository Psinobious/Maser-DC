package User

import(
	"net/http"
)
type createUser struct {
	token string
	firstName string
	lastName string
	email string	
	password string
}
func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
	
}