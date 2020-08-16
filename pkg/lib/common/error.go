package common

import "errors"

var (
	// Parameter errors
	ErrInvalidIDParameter        = errors.New(InvalidIDParameterMsg)
	ErrInvalidLowOffsetParameter = errors.New(InvalidLowOffsetMsg)
	ErrInvalidLimitParameter     = errors.New(InvalidLimitParameterMsg)
	ErrInvalidOrderByParameter   = errors.New(InvalidOrderByParameterMsg)
	ErrInvalidOrderTypeParameter = errors.New(InvalidOrderTypeParameterMsg)
)
