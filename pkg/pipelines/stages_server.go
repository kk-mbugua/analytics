package pipelines

import (
	"context"
	"log"
	"main/pkg/proto/pb"
)

type PipelineStageServer struct {
	pb.UnimplementedPipelineStageServiceServer
	PipelineStageStore *DatabasePipelineStageStore
}

func NewPipelineStageServer(PipelineStageStore *DatabasePipelineStageStore) *PipelineStageServer {
	return &PipelineStageServer{
		UnimplementedPipelineStageServiceServer: pb.UnimplementedPipelineStageServiceServer{},
		PipelineStageStore:                      PipelineStageStore,
	}
}

// CreatePipelineStage creates a new pipeline stage
func (s *PipelineStageServer) CreatePipelineStage(ctx context.Context, req *pb.CreatePipelineStageRequest) (*pb.CreatePipelineStageResponse, error) {
	log.Println("CreatePipelineStage")

	stage := PbStageToStage(req.Stage)

	return &pb.CreatePipelineStageResponse{Stage: StageToPbStage(stage)}, nil
}

// GetPipelineStage gets a pipeline stage by id
func (s *PipelineStageServer) GetPipelineStage(ctx context.Context, req *pb.GetPipelineStageRequest) (*pb.GetPipelineStageResponse, error) {
	log.Println("GetPipelineStage")

	stage, err := s.PipelineStageStore.ReadStage(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPipelineStageResponse{Stage: StageToPbStage(stage)}, nil
}

// GetPipelineStages gets all pipeline stages for a pipeline
func (s *PipelineStageServer) GetPipelineStages(ctx context.Context, req *pb.GetPipelineStagesRequest) (*pb.GetPipelineStagesResponse, error) {
	log.Println("GetPipelineStages")

	stages, err := s.PipelineStageStore.ReadStages(req.PipelineId)
	if err != nil {
		return nil, err
	}

	return &pb.GetPipelineStagesResponse{Stages: StagesToPbStages(stages)}, nil
}

// UpdatePipelineStage updates a pipeline stage
func (s *PipelineStageServer) UpdatePipelineStage(ctx context.Context, req *pb.UpdatePipelineStageRequest) (*pb.UpdatePipelineStageResponse, error) {
	log.Println("UpdatePipelineStage")

	stage := PbStageToStage(req.Stage)

	return &pb.UpdatePipelineStageResponse{Stage: StageToPbStage(stage)}, nil
}

// DeletePipelineStage deletes a pipeline stage by id
func (s *PipelineStageServer) DeletePipelineStage(ctx context.Context, req *pb.DeletePipelineStageRequest) (*pb.DeletePipelineStageResponse, error) {
	log.Println("DeletePipelineStage")

	err := s.PipelineStageStore.DeleteStage(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePipelineStageResponse{
		Status: &pb.StatusP{
			Code:    200,
			Message: "Pipeline stage deleted",
		},
	}, nil
}
