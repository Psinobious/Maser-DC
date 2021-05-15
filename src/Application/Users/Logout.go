package User

import(
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"OAuth"
	"context"
)

type LogOut struct {
	AccessToken string
}
func (u *UserHandler) Logout(w http.ResponseWriter, r *http.Request){
	var conn *grpc.ClientConn

	c, err := r.Cookie("Access_Token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// For any other type of error, return a bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Access_Token{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	message := OAuth.LogoutToken{
		AccessToken: claims.accessToken,
	}
	conn, err = grpc.Dial(u.address, grpc.WithInsecure()) 
	auth := OAuth.NewOAuthClient(conn)
	_, err = auth.InvalidateTokens(context.Background(), &message)

}