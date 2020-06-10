package grpc

import (
	"context"
	"fmt"
)

type allowAuthenticator struct{}

func (a allowAuthenticator) Authenticate(ctx context.Context) error {

	fmt.Println("allow",ctx)
	return nil
}

// AllowAuthenticator is an implementation of Authenticator that simply
// always returns success. This implementation can be used in case a
// gRPC server needs to be started that does not perform any
// authentication (e.g., one listening on a UNIX socket with restricted
// file permissions).
var AllowAuthenticator Authenticator = allowAuthenticator{}
