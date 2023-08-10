package jwt

import (
	"ex-server/internal/auth/entity"
	"ex-server/internal/auth/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// TODO: move to config
var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func GenerateJWT(id, login string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Id:    id,
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func ReadToken(signedToken string) (*entity.User, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return nil, exception.ErrTokenInvalid
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, exception.ErrTokenExpired
	}
	return &entity.User{
		Id:    claims.Id,
		Login: claims.Login,
	}, nil
}

func GenerateRefresh() (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	return refreshToken.SignedString(jwtKey)
}

func ValidateRefresh(refresh string) error {
	token, err := jwt.Parse(refresh, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.ErrTokenInvalid
		}

		return jwtKey, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if int(claims["sub"].(float64)) == 1 {
			return nil
		}
	}

	return exception.ErrTokenInvalid
}
