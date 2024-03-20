package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("in middleware")
		tokenString := r.Header.Get("Authorization")
		t := strings.Replace(tokenString, "Bearer ", "", -1)
		verifiedToken, err := VerifyToken(t)
		if err != nil {
			fmt.Println("error in token verification is ", err)
		} else {
			data := verifiedToken.Claims.(jwt.MapClaims)
			fmt.Println(data)
		}
		next.ServeHTTP(w, r)
	})
}
