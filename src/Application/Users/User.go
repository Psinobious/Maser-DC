package User

import (
	"time"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"
	"OAuth"
)
type UserHandler struct{
	UserRepository UserRepository
	OAuth OAuth.OAuthClient
	address string
	connection grpc.ClientConn
}
type UserRepository interface {
	updateUser(UserID, FirstName, LastName, Email, status string)(error)
	createUser(UserID, FirstName, LastName, Email, Password string)(error)
	deleteUser(UserID string)(error)
	changePassword(UserID, password string)(error)
	findUser(UserID string) (User, error)
}
type UserNeo4jRepository struct {
	Driver neo4j.Driver
}
type User struct {
	userID string
	alias string
	firstName string
	lastName string
	email string	
	password string
	status string
	salt []byte
	date_created time.Time
	update_time time.Time
}

func (u *UserNeo4jRepository) findUser(UserID string) (User, error){
	session := u.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	query := "Match (u:USER{UserID:'$UserID') return n"
	parameters := map[string] interface{}{
		"UserID": UserID,
	}
	result, err := session.ReadTransaction(func(tx neo4j.Transaction)(interface{}, error){
		result, err := tx.Run(query, parameters)
		return result.Record(), err
	})
	if err != nil {
		panic(err)
	}	
	return result.(User), nil	
}
func (u *UserNeo4jRepository) persistUser(UserID, FirstName, LastName, Email, Password string, tx neo4j.Transaction) error{
	query := `CREATE(n:USER{UserID:$UserID, FirstName:$FirstName, 
				LastName:$LastName, Email:$Email, Password:$Password, 
				Date_Created:TIMESTAMP(), Date_Last_Modified: TIMESTAMP()
				})`
				
	parameters := map[string]interface{}{
		"UserID": UserID,
		"FirstName": FirstName,
		"LastName": LastName,
		"Email": Email,
		"Password": Password,
	}
	_, err := tx.Run(query, parameters)
	if err != nil {
		panic(err)
	}
	return nil
}
func (u *UserNeo4jRepository) deleteUser(UserID string, tx neo4j.Transaction) {
	query := "MATCH(n:USER{UserID:$UserID}) DELETE n"
	parameters := map[string]interface{}{
			"UserID": UserID,
		}
	_, err := tx.Run(query, parameters)
	if err != nil {
		panic(err)
	}
}
func (u *UserNeo4jRepository) updateUser(UserID, FirstName, LastName, Email, Status string) error{
	session := u.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})

	defer session.Close()

	query := `MATCH (n:USER {UserID: $UserID})
				SET n.FirstName = $FirstName, n.LastName = $LastName, n.Status = $Status`
	parameters := map[string]interface{}{		
		"UserID": UserID,
		"FirstName": FirstName,
		"LastName": LastName,
		"Status": Status,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (u *UserNeo4jRepository) changePassword(UserID, Password string) error{
	session := u.Driver.NewSession(neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeWrite,
	})
	query :="MATCH(n:USER{UserID:$UserID}) SET n.Password = $Password"
	parameters := map[string]interface{}{
		"UserID": UserID,
		"Password": Password,
	}
	_, err := session.WriteTransaction(func(tx neo4j.Transaction)(interface{}, error){
		_, err := tx.Run(query, parameters)
		return nil, err
	})
	if err != nil {
		panic(err)
	}
	return nil
}
func (u *UserNeo4jRepository) changeEmail(UserID, Email string)error{
	return nil
}

