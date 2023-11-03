package validator

import "errors"

var (
	ErrInvalidEmail    = errors.New("Invalid email address")
	ErrStringLengthMin = errors.New("String length is less than required")
	ErrStringLengthMax = errors.New("String length is exceeded than required")
	ErrMissingField    = errors.New("Field can't be blank")
	ErrValueOutOfRange = errors.New("Value out of range")
)
