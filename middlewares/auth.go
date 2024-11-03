package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aphrem-thomas/password-manager/utils"
	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		userName := ""
		fmt.Println("in middleware request -->", r)
		tokenString := r.Header.Get("Authorization")
		fmt.Println("token string", tokenString)
		t := strings.Replace(tokenString, "Bearer ", "", -1)
		verifiedToken, err := VerifyToken(t)
		if err != nil {
			fmt.Println("error in token verification is ", err)
		} else {
			data := verifiedToken.Claims.(*UserClaims)
			userName = data.Username
			fmt.Println(data.Username)

		}
		ctx = context.WithValue(r.Context(), utils.UserName, userName)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
