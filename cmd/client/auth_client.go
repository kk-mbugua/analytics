package authclient

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"main/pkg/proto/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type AuthClient struct {
	Client pb.AuthServiceClient
}

func NewAuthClient(serverAddr, certFile, keyFile, caFile string) (*AuthClient, error) {
	// Load client certificate and key
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Load CA certificate
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create TLS configuration with client certificate and CA certificate
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}

	// Dial with TLS credentials
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))
	if err != nil {
		return nil, err
	}

	client := pb.NewAuthServiceClient(conn)

	return &AuthClient{
		Client: client,
	}, nil
}
