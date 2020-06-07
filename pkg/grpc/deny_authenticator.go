package grpc

import (
	"context"
        "fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
        "google.golang.org/grpc/metadata"
)

type denyAuthenticator struct {
	err error
}

// NewDenyAuthenticator creates an Authenticator that always returns an
// UNAUTHENTICATED error with a fixed error message string. This
// implementation can be used in case a gRPC server needs to be
// administratively disabled without shutting it down entirely.
func NewDenyAuthenticator(message string) Authenticator {
	return &denyAuthenticator{
		err: status.Error(codes.Unauthenticated, message),
	}
}

func (a denyAuthenticator) Authenticate(ctx context.Context) error {

	md, ok := metadata.FromIncomingContext(ctx)
	//.Get("authorization")
	fmt.Println(ok)
	fmt.Println("deny Authenticate",ctx,md)
          //return nil
	  return a.err
}
