package main

import (
	"fmt"
	"log"
	"main/cmd/server/config"
	"main/pkg/auth"
	"main/pkg/db"
	pipelines "main/pkg/pipelines"
	"main/pkg/proto/pb"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

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
		creds, err := credentials.NewServerTLSFromFile("certs/dev/server.crt", "certs/dev/server.key")
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		log.Printf("Loaded insecure certificates %v", insecure.NewCredentials())
		return creds
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
	leadStore := pipelines.NewDatabaseLeadStore(db.DB)

	pipelinesServer := pipelines.NewPipelineServer(pipelineStore)
	pipelineStagesServer := pipelines.NewPipelineStageServer(stageStore)
	pipelineStageLabelsServer := pipelines.NewStageLabelServer(stageLabelStore)
	leadServer := pipelines.NewLeadServer(leadStore)

	interceptor := auth.NewAuthInterceptor()

	grpcServer := grpc.NewServer(
		grpc.Creds(certificates),
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)
	reflection.Register(grpcServer)

	pb.RegisterPipelineServiceServer(grpcServer, pipelinesServer)
	pb.RegisterPipelineStageServiceServer(grpcServer, pipelineStagesServer)
	pb.RegisterStageLabelServiceServer(grpcServer, pipelineStageLabelsServer)
	pb.RegisterLeadServiceServer(grpcServer, leadServer)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", configurations.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server started on port %v\n", configurations.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
