package main

import (
	"context"
	"fmt"
	"log"
	"main/cmd/server/config"
	"main/pkg/db"
	pipelines "main/pkg/pipelines"
	"main/pkg/proto/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func UnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func StreamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	log.Println("--> stream interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

func loadCertifcate(environment string) credentials.TransportCredentials {
	switch environment {
	case "production":
		creds, err := credentials.NewServerTLSFromFile("certs/prod/cert.pem", "certs/prod/key.pem")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		log.Printf("Loaded production certificates %v", creds)
		return creds
	case "staging":
		creds, err := credentials.NewServerTLSFromFile("certs/staging/cert.pem", "certs/staging/key.pem")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		log.Printf("Loaded staging certificates %v", creds)
		return creds
	case "development":
		log.Printf("Loaded insecure certificates %v", insecure.NewCredentials())
		return insecure.NewCredentials()
	default:
		log.Printf("Loaded insecure certificates %v", insecure.NewCredentials())
		return insecure.NewCredentials()
	}
}

func main() {
	configurations, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db.InitDB(
		configurations.PostgresHost,
		configurations.PostgresPort,
		configurations.PostgresDB,
		configurations.PostgresUser,
		configurations.PostgresPass,
	)

	log.Printf("Starting server on port %v\n", configurations.Port)

	certificates := loadCertifcate(configurations.Environment)

	pipelineStore := pipelines.NewDatabasePipelineStore(db.DB)
	stageStore := pipelines.NewDatabasePipelineStageStore(db.DB)
	stageLabelStore := pipelines.NewDatabaseStageLabelStore(db.DB)

	pipelinesServer := pipelines.NewPipelineServer(pipelineStore)
	pipelineStagesServer := pipelines.NewPipelineStageServer(stageStore)
	pipelineStageLabelsServer := pipelines.NewStageLabelServer(stageLabelStore)
	interceptor := pipelines.NewAuthInterceptor()

	grpcServer := grpc.NewServer(
		grpc.Creds(certificates),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	reflection.Register(grpcServer)

	pb.RegisterPipelineServiceServer(grpcServer, pipelinesServer)
	pb.RegisterPipelineStageServiceServer(grpcServer, pipelineStagesServer)
	pb.RegisterStageLabelServiceServer(grpcServer, pipelineStageLabelsServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configurations.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server started on port %v\n", configurations.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
