package service

import "errors"

// ErrInvalidCEP é retornado quando o CEP fornecido é inválido.
var ErrInvalidCEP = errors.New("CEP inválido")
