package utils

import (
	"errors"
	"go-grpc-auth-svc/pkg/models"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

func (wrapper *JwtWrapper) GenerateToken(user models.User) (signedToken string, err error) {
	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(wrapper.ExpirationHours)).Unix(),
			Issuer:    wrapper.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(wrapper.SecretKey))

	if err != nil {
		return "", jwt.ErrECDSAVerification
	}

	return signedToken, nil
}

func (wrapper *JwtWrapper) ValidToken(signedToken string) (claims *jwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(wrapper.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("could not parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
