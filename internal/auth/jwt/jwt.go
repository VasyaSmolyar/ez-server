package jwt

import (
	"ex-server/internal/auth/entity"
	"ex-server/internal/auth/exception"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func Init(cfg *viper.Viper) *JWTHelper {
	return &JWTHelper{
		jwtKey:             []byte(cfg.GetString("JWT.Secret")),
		accessLiveSeconds:  cfg.GetInt("JWT.AccessLiveSeconds"),
		refreshLiveSeconds: cfg.GetInt("JWT.RefreshLiveSeconds"),
	}
}

type JWTHelper struct {
	jwtKey             []byte
	accessLiveSeconds  int
	refreshLiveSeconds int
}

type JWTClaim struct {
	Id    string `json:"id"`
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func (h *JWTHelper) GenerateJWT(id, login string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(h.accessLiveSeconds) * time.Second)
	claims := &JWTClaim{
		Id:    id,
		Login: login,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(h.jwtKey)
}
func (h *JWTHelper) ReadToken(signedToken string) (*entity.User, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return h.jwtKey, nil
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

func (h *JWTHelper) GenerateRefresh() (string, error) {
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Duration(h.refreshLiveSeconds) * time.Second).Unix()

	return refreshToken.SignedString(h.jwtKey)
}

func (h *JWTHelper) ValidateRefresh(refresh string) error {
	token, err := jwt.Parse(refresh, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, exception.ErrTokenInvalid
		}

		return h.jwtKey, nil
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
