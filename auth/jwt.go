package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

// here we are declaring a secret key that will be used later for generating JWT
// for now the key is secretkey
var jwtKey = []byte("secretkey")

// Define a custom struct for JWT claims which will ultimately become the payload of the JWT
type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// this function will return the generated JWT string. Here we set default expiration time for 1 hour
func GenerateJWT(email string, username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

// Take in the token string coming from the client HTTP request header and then validate it.
// here we will try to parse the JWT into claims using the JWT package’s helper method “ParseWithClaims”.
func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("could not parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return
}
