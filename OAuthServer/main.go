package main

import (
	"context"
	"time"
	"google.golang.org/grpc"
	"github.com/satori/go.uuid"
	"github.com/go-redis/redis/v8"
	"net"
	"log"
	oauth "OAuth"
)

var ctx = context.Background()


type AccessTokens struct {
	accessToken string
	tokenType string
	expires_in int
	refreshToken string
	scope string
}
type SuccessfulResponse struct {
	clientID string
	scope string
	tokenType string
}
type RefreshRequest struct {
	grantType string
	refreshToken string
	clientID string
	clientSecret string
}
type AuthorizationGrant struct {
	username string
	password string
	grantType string
}
type Server struct {
	oauth.UnimplementedOAuthServer
	rdb *redis.Client
}
const (
	port = ":50051"
)

func main(){

	_ = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
	})	
	lis, err := net.Listen("tcp", port)

	if err != nil{
		panic(err)
	}
	s := grpc.NewServer()
	oauth.RegisterOAuthServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
func (s *Server) OAuthTest(ctx context.Context, in *oauth.TestMessage)(*oauth.Response, error){
	return &oauth.Response{
		Code: "Good",
	}, nil
}
func (s *Server) RequestAccessTokens(ctx context.Context, in *oauth.AuthorizationGrant)(*oauth.AccessTokens, error){
	accessToken := uuid.NewV4().String()
	refreshToken := uuid.NewV4().String()

	expires_in := 20*time.Minute
	successfulResponse := SuccessfulResponse{in.GetClientID(), in.GetScope(), in.GetGrant_Type()}
	err := s.rdb.Set(ctx, accessToken, successfulResponse, expires_in).Err()
	err = s.rdb.Set(ctx, refreshToken, in.GetClientID(), time.Hour).Err()

    if err != nil {
        panic(err)
	}
	return &oauth.AccessTokens{
		AccessToken: accessToken,
		TokenType: in.GetGrant_Type(),
		Expires_In: int64(expires_in),
		RefreshToken: refreshToken,
		Scope: in.GetScope(),
	}, nil
}
func(s *Server) RefreshAccessTokens(ctx context.Context, in *oauth.AccessTokenRequest)(*oauth.AccessTokens, error){
	accessToken := uuid.NewV4().String()
	
	expires_in := 20*time.Minute
	successfulResponse := AccessTokens{accessToken, in.GetGrant_Type(), 3600, in.GetRefreshToken(), in.GetScope()}

	err := s.rdb.Set(ctx, accessToken, successfulResponse, expires_in).Err()
	if err != nil {
        panic(err)
	}

	return &oauth.AccessTokens{
		AccessToken: accessToken,
		TokenType: in.GetGrant_Type(),
		Expires_In: int64(expires_in),
		RefreshToken: in.GetRefreshToken(),
		Scope: in.Scope,
	}, nil
}
func(s *Server) InvalidateTokens(ctx context.Context, in *oauth.LogoutToken)(*oauth.Response, error){
	
	return nil, nil
}