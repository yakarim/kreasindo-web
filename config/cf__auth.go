package config

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/savsgio/go-logger"
)

var jwtSignKey = []byte("bismillah")

// UserCredential ...
type UserCredential struct {
	Username []byte `json:"username"`
	Password []byte `json:"password"`
	jwt.StandardClaims
}

// GenerateToken ...
func (c Config) GenerateToken(username []byte, password []byte) (string, time.Time) {
	logger.Debugf("Create new token for user %s", username)

	expireAt := time.Now().Add(120 * time.Minute)

	// Embed User information to `token`
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS512, &UserCredential{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	})

	// token -> string. Only server knows the secret.
	tokenString, err := newToken.SignedString(jwtSignKey)
	if err != nil {
		logger.Error(err)
	}

	return tokenString, expireAt
}

func validateToken(requestToken string) (*jwt.Token, *UserCredential, error) {
	logger.Debug("Validating token...")

	user := &UserCredential{}
	token, err := jwt.ParseWithClaims(requestToken, user, func(token *jwt.Token) (interface{}, error) {
		return jwtSignKey, nil
	})

	return token, user, err
}
