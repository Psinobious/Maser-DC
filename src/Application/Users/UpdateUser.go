package User

import(
	"net/http"
	
)
type updateUser struct{
	token string
	firstName string
	lastName string
	status string	
}
func (u *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request){

}