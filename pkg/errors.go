package pkg

import "errors"

var (
	ErrBadWrite                = errors.New("bytes written out is not equal to bytes received")
	ErrServiceNameFlagRequired = errors.New("the service name flag must be passed to protoc as `svc=<name>`")
)
