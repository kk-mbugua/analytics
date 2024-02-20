package authclient

import (
	"fmt"
	"main/pkg/proto/pb"

	"google.golang.org/grpc"
)

type AuthClient struct {
	Client pb.AuthServiceClient
}

func NewAuthClient(serverAddr string) (*AuthClient, error) {
	// creds, err := credentials.NewServerTLSFromFile("certs/dev/server.crt", "certs/dev/server.key")
	// if err != nil {
	// 	return nil, fmt.Errorf("could not load tls cert: %s", err)
	// }
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("did not connect: %v", err)
	}

	client := pb.NewAuthServiceClient(conn)

	return &AuthClient{
		Client: client,
	}, nil
}
