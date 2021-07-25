package auth

import "github.com/dgrijalva/jwt-go"

type Service interface {
	GenerateToken(userID int) (string, error)
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
