package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}
type jwtservice struct {
}

var sc = []byte("wkwk")

func NewService() *jwtservice {
	return &jwtservice{}
}

func (s *jwtservice) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(sc)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtservice) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(sc), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
