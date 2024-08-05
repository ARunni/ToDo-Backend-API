package helper

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type AuthUserClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func GetTokenFromHeader(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return header
}

func ExtractUserIDFromToken(tokenString string) (int, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AuthUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte("user@todo"), nil
	})

	if err != nil {
		fmt.Println("errors:-", err)
		return 0, "", err
	}
	claims, ok := token.Claims.(*AuthUserClaims)
	if !ok {
		return 0, "", fmt.Errorf("invalid token claims")
	}
	return int(claims.Id), claims.Name, nil

}
