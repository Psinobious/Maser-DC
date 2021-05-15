package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"time"
	"log"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"Infrastructure"
)
var(
	dbUri = "bolt://localhost:7687"
	OAuthPort = "localhost:50051"
)
func main() {
//	userRepository := UserNeo4jRepository{
//		Driver: driver(dbUri,neo4j.BasicAuth("neo4j","orchestrator","")),
//	}
//	Handler := &WebServiceHandler{
//		UserRepository: &userRepository,
//	}
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET","POST","PUT","DELETE", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	
	var Handler Infrastructure.WebServiceHandler
	Handler.OAuthServerAddress = OAuthPort

	r.HandleFunc("/Login", Handler.UserRepository.Login).Methods("GET")
	r.HandleFunc("/Logout", Handler.UserRepository.Logout).Methods("POST")
	r.HandleFunc("/UpdateUser", Handler.UserRepository.UpdateUser).Methods("POST")
	r.HandleFunc("/Refresh", Handler.UserRepository.Refresh).Methods("GET")
	r.HandleFunc("/code/{id}", Handler.UserRepository.CreateUser).Methods("POST")

	r.HandleFunc("/Test", Handler.Test.TestOAuthConnection).Methods("GET")
	
	s := &http.Server{
		Addr:           ":8082",
		Handler:        handlers.CORS(headers, methods, origins)(r),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Starting server on port 8082.....")
	log.Fatal(s.ListenAndServe())

}
func driver(target string, token neo4j.AuthToken) neo4j.Driver {
	result, err := neo4j.NewDriver(target, token)
	if err != nil {
		panic(err)
	}
	return result
}
func handleError(err error) {
	if err != nil {
		panic(err)
	}
}