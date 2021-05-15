package Test

import(
	"net/http"
	//"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"OAuth"
	"context"
)

type TestHandler struct {
	address string
	connection grpc.ClientConn
}
const (
	port = "localhost:50051"
)
func (t *TestHandler) TestOAuthConnection(w http.ResponseWriter, r *http.Request){
	fmt.Println("Running Test method")
	var conn *grpc.ClientConn
	
	conn, err := grpc.Dial(port, grpc.WithInsecure()) 
	defer conn.Close()	
	auth := OAuth.NewOAuthClient(conn)
	message := OAuth.TestMessage{
		Message: "test",
	}
	response, err := auth.OAuthTest(context.Background(), &message)
	if(err != nil){
		panic(err)
	}
	fmt.Fprintf(w, response.Code, r.URL.Path[1:])
}