package main

import (
	"context"
	"fmt"
	"log"
	"main/cmd/server/config"
	"main/pkg/db"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// Initialize the database
	db.InitDB(c.PostgresHost, c.PostgresPort, c.PostgresDB, c.PostgresUser, c.PostgresPass)
	log.Printf("Starting server on port %d\n", c.Port)
	// TODO: Create Store instaces for services

	// TODO: Create Server instances for services

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	reflection.Register(grpcServer)

	// TODO: Register Servers to pb.Register<Package_Service>Server(grpcServer, Server)

	fmt.Println("TCP PORT:", c.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", c.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server started on port %d\n", c.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
