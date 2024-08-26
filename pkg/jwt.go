package pkg

import (
	"time"

	"github.com/MXkodo/verif/config"
	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID string, ip string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"ip":      ip,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(config.JWTSecret))
}
