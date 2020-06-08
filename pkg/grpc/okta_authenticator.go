package grpc

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
	fmt.Println("md = ", md.Get("Authorization"))
	//return nil

	rawToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return a.err
	}
	claims := &jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(rawToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("#secret#"), nil
	})
	if err != nil {
		return a.err
	}
	if !token.Valid {
	   return a.err
	} else {
	  return nil
	}
}
