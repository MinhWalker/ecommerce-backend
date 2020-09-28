package security

import (
	"ecommerce-backend/model"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

const JWT_KEY = "hhhgfdshgfhsdgfshjgfshjdgf"

func GenToken(user model.User) (string, error) {
	jwtExpConfig := os.Getenv("JwtExpires")

	jwtExpValue, _ := strconv.Atoi(jwtExpConfig) // Convert to int

	jwtExpDuration := time.Hour * time.Duration(jwtExpValue) // Convert to hour

	claims := &model.JwtCustomClaims{
		UserId: user.UserID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
