package pipelines

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// AuthInterceptor is an interceptor for Authentication and Authorization.
type AuthInterceptor struct {
}

// NewAuthInterceptor creates a new instance of AuthInterceptor.
func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

type key int

const (
	UserID key = iota
	OrganisationID
	Authorization
	RequestAuth
	BranchID
)

// Unary returns a server interceptor for authentication and authorization.
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)

		// TODO: implement authentication and authorization
		return handler(ctx, req)
	}
}

// Stream returns a server interceptor for streaming calls for authentication and authorization.
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {

		return handler(srv, stream)
	}
}
