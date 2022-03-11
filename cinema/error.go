package main

import "errors"

var (
	ErrAlreadyExists = errors.New("cinema already exists")
	ErrNotFound      = errors.New("not found")
	ErrBadRouting    = errors.New("invalid routing")
)
