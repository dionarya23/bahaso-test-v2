package userusecase

import "errors"

var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyUsed = errors.New("email already used")
var ErrInvalidPassword = errors.New("wrong password")
var ErrInvalidToken = errors.New("token not valid")
var ErrExpiredToken = errors.New("token expired")
var ErrTokenNotFound = errors.New("token not found")
var ErrInvalidUser = errors.New("wrong email or password")
