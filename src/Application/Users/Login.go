package User

import(
	"net/http"
	"encoding/json"
	"Password"
	"google.golang.org/grpc"
	"context"
	"OAuth"
	"time"
	"github.com/dgrijalva/jwt-go"
)
var jwtKey = []byte("my_secret_key")
type UserLogin struct{
	userID string	`json:"username"`
	password string	`json:"password"`
}

type OAuthClient struct {
	address string	
}
type Access_Token struct {
	accessToken string
    tokenType string
    scope string
	jwt.StandardClaims
}
type Refresh_Token struct{
	refreshToken string
	tokenType string
	scope string
	jwt.StandardClaims
}
func (u *UserHandler) Login(w http.ResponseWriter, r *http.Request){
	var conn *grpc.ClientConn
	var Request UserLogin
	var handler Password.PasswordHandler
	json.NewDecoder(r.Body).Decode(&Request)
	
	//Checks to see if user exist
	user, err := u.UserRepository.findUser(Request.userID)
	if(err != nil){
		panic(err)
	}
	//Verify user is legitimate via password check
	sentinel := handler.VerifyPassword(user.password, Request.password, user.salt)	
	
	if(sentinel == true){
		//Sends authorization grant to the authorization server
		conn, err = grpc.Dial(u.address, grpc.WithInsecure()) 
		defer conn.Close()		

		auth := OAuth.NewOAuthClient(conn)
		message := OAuth.AuthorizationGrant {
			Grant_Type: "authentication",
			ClientID: user.userID,
			Scope: "regular",
		}
		//Recieve the access and refresh tokens
		response, err := auth.RequestAccessTokens(context.Background(), &message)
		
		if(err != nil){
			panic(err)
		}

		//Create cookie for Access token
		expirationTime := time.Now().Add(time.Duration(response.Expires_In)*time.Minute)
		
		accessToken := &Access_Token{
			accessToken: response.AccessToken,
    		tokenType: response.TokenType,
			scope: response.Scope,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessToken)
		tokenString, err := token.SignedString(jwtKey)
		if(err != nil){
			panic(err)
		}
		a := http.Cookie{
			Name:    "Access_Token",
			Value:   tokenString,
			Expires: expirationTime,
			HttpOnly: false,
		}
		
		http.SetCookie(w, &a)
		
		//Create cookie for refresh token

		refreshToken := &Refresh_Token{
			refreshToken: response.RefreshToken,
			tokenType: response.TokenType,
			scope: response.Scope,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},		
		}
		token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshToken)
		tokenString, err = token.SignedString(jwtKey)
		if(err != nil){
			panic(err)
		}
		r := http.Cookie{
			Name: "Refresh_Token", 
			Value: tokenString,
			Expires: time.Now().Add(time.Hour),
			HttpOnly: false,
		}
		//Send Cookies
		
		http.SetCookie(w, &r)
	}else{
		
	}	
}
