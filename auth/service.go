package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(encToken string) (*jwt.Token, error)
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

func (j *jwtService) ValidateToken(encToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}
	return token, nil
}
