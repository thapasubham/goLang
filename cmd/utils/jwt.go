package utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {

	timeDur, _ := strconv.Atoi(LoadEnv("JWT_EXPIRATION"))
	fmt.Println(timeDur)
	expiration := time.Second * time.Duration(timeDur)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":  userID,
			"exp": time.Now().Add(expiration).Unix(),
		})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetIDJwt(r *http.Request) (int, error) {
	tokenString := r.Header.Get("token")

	if tokenString == "" {
		return 0, fmt.Errorf("no token found")
	}

	token, err := validateToken(tokenString)

	if err != nil {
		fmt.Println(err)
		return 0, nil
	}

	if !token.Valid {
		return 0, fmt.Errorf("Token no valid")
	}
	claims := token.Claims.(jwt.MapClaims)
	value := claims["id"].(float64)

	fmt.Println(value)
	user_id := int(value)

	return user_id, nil

}

func validateToken(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["args"])
		}
		return []byte(LoadEnv("SECRET")), nil
	})

}
