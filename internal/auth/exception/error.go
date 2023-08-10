package exception

import "errors"

var ErrWrongConfig error = errors.New("wrong config")
var ErrWrongCreds error = errors.New("wrong credentials")
var ErrAlreadyCreated error = errors.New("the user has already registered")
var ErrHashingFailed error = errors.New("hashing failed")
var ErrTokenExpired error = errors.New("token expired")
var ErrTokenInvalid error = errors.New("token invalid")
