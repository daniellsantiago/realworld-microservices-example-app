package pkg

import "errors"

var (
	ErrRecordNotFound      = errors.New("record not found")
	ErrInternal            = errors.New("an internal error has occurred")
	ErrValidation          = errors.New("invalid state")
	ErrUnProcessableEntity = errors.New("unprocessable entity")
)
