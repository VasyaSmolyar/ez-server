package request

import (
	"github.com/golang-jwt/jwt/v5"
)

type Credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserClaim struct {
	Id    string `json:"id"`
	Login string `json:"title"`
	Exp   int64  `json:"exp"`
	jwt.RegisteredClaims
}

type RefreshClaim struct {
	Sub int   `json:"sub"`
	Exp int64 `json:"exp"`
	jwt.RegisteredClaims
}
