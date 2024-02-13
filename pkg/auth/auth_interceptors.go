package auth

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

type key int

const (
	UserID key = iota
	OrganisationID
	Authorization
	RequestAuth
	BranchID
)

type AuthInterceptor struct {
}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Println("--> unary interceptor: ", info.FullMethod)
		requestMetadata, err := AuthRequest(ctx, "accessToken", info.FullMethod)
		if err != nil {
			return nil, err
		}
		ctx = context.WithValue(ctx, UserID, requestMetadata.UserID)
		ctx = context.WithValue(ctx, OrganisationID, requestMetadata.OrganisationID)
		ctx = context.WithValue(ctx, Authorization, requestMetadata.Authorization)
		ctx = context.WithValue(ctx, RequestAuth, requestMetadata.RequestAuth)
		ctx = context.WithValue(ctx, BranchID, requestMetadata.BranchID)
		return handler(ctx, req)
	}
}
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
