package grpc

import (
	"context"
        "fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
        "google.golang.org/grpc/metadata"
)

type oktaAuthenticator struct {
	err error
}

// NewOktaAuthenticator creates an Authenticator that for okta
func NewOktaAuthenticator(message string) Authenticator {
	return &oktaAuthenticator{
		err: status.Error(codes.Unauthenticated, message),
	}
}

func (a oktaAuthenticator) Authenticate(ctx context.Context) error {

	md, ok := metadata.FromIncomingContext(ctx)
	//.Get("authorization")
	fmt.Println(ok)
	fmt.Println("okta Authenticate",ctx,md)
          //return nil
	  return a.err
}
