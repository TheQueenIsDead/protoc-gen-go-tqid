package pkg

import "errors"

var (
	ErrBadWrite = errors.New("bytes written out is not equal to bytes received")
)
