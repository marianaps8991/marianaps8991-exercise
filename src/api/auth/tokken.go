package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"api/utils/console"
	"config"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userId uint32, username string, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30)
	tokken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokken.SignedString(config.SECRET)
}

func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.SECRET, nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		console.DisplayData(claims)
	}
	return nil

}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}

	return ""
}

func ExtractTokenID(r *http.Request) (uint32, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.SECRET, nil
	})

	if err != nil {

		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		uid, err := strconv.ParseUint(
			fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(uid), nil
	}
	return 0, err
}

func AdminValid(r *http.Request) (bool, error) {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return config.SECRET, nil
	})

	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {

		isAdmin := claims["role"] == "admin"
		if isAdmin {
			return true, nil
		}

	}
	return false, err

}
