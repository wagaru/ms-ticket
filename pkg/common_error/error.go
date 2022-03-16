package common_error

import "errors"

var (
	ErrAlreadyExists = errors.New("data already exists")
	ErrNotFound      = errors.New("data not found")

	ErrBadRouting   = errors.New("invalid routing")
	ErrInvalidInput = errors.New("invalid request data")
)
