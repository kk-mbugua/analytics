package pipelines

import (
	"context"
	"log"
	utils "main/pkg/Util"
	"main/pkg/proto/pb"

	"github.com/google/uuid"
)

type PipelineServer struct {
	pb.UnimplementedPipelineServiceServer
	PipelineStore *DatabasePipelineStore
}

func NewPipelineServer(pipelineStore *DatabasePipelineStore) *PipelineServer {
	return &PipelineServer{
		UnimplementedPipelineServiceServer: pb.UnimplementedPipelineServiceServer{},
		PipelineStore:                      pipelineStore,
	}
}

func (s *PipelineServer) CreatePipeline(ctx context.Context, req *pb.CreatePipelineRequest) (*pb.CreatePipelineResponse, error) {
	log.Println("CreatePipeline")

	pipeline := PbPipelineToPipeline(req.Pipeline)
	pipeline.ID = uuid.New().String()
	pipeline.SerialNumber, _ = utils.CreateSerialNumber("PIP", pipeline.ID)

	// TODO: validate pipeline
	// TODO: check if pipeline name already exists
	return &pb.CreatePipelineResponse{}, nil
}
