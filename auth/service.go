package auth

import (
	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct{}

var SECRET_KEY = []byte("DNUF")

func NewJwtService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}
