package utils

import "errors"

var (
	ErrPortRequired = errors.New("PORT environment variable is required")
	ErrInvalidPort  = errors.New("invalid PORT value")
)
