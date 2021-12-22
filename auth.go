package main

import(
	"time"
	"net/http"
	// "fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
)
var jwtKey = []byte("oqdaryoservis")


var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32),
)

func CreateToken(userName string, w http.ResponseWriter) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims {
			Issuer: userName,
			ExpiresAt: time.Now().AddDate(0,6,0).Unix(),
		})
	if tokenString, err := token.SignedString(jwtKey); err== nil {
		value := map[string]string{
			"tkn": tokenString,
		}
		encoded, err := cookieHandler.Encode("token", value);
		if err == nil {
			cookie := &http.Cookie{
				Name:  "token",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(w, cookie)
		}
	}
}


func CheckToken(r *http.Request) bool {
	var token string
	if cookie, err := r.Cookie("token"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("token", cookie.Value, &cookieValue); err == nil {
			token = cookieValue["tkn"]
		}
	}

	if token!=""{

		if tkn, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {return jwtKey, nil});err == nil {
			return tkn.Valid
		}
	}
	return false

}
