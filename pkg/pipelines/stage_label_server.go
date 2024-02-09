package pipelines

import (
	"context"
	"log"
	"main/pkg/proto/pb"
)

type StageLabelServer struct {
	pb.UnimplementedStageLabelServiceServer
	StageLabelsStore *DatabaseStageLabelStore
}

func NewStageLabelServer(stageLabelsStore *DatabaseStageLabelStore) *StageLabelServer {
	return &StageLabelServer{
		StageLabelsStore: stageLabelsStore,
	}
}

// CreateStageLabel creates a new stage label
func (s *StageLabelServer) CreateStageLabel(ctx context.Context, req *pb.CreateStageLabelRequest) (*pb.CreateStageLabelResponse, error) {
	log.Println("CreateStageLabel")

	stageLabel := PbStageLabelToStageLabel(req.StageLabel)

	return &pb.CreateStageLabelResponse{StageLabel: StageLabelToPbStageLabel(stageLabel)}, nil
}

// GetStageLabel gets a stage label by id
func (s *StageLabelServer) GetStageLabel(ctx context.Context, req *pb.GetStageLabelRequest) (*pb.GetStageLabelResponse, error) {
	log.Println("GetStageLabel")

	stageLabel, err := s.StageLabelsStore.ReadStageLabel(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetStageLabelResponse{StageLabel: StageLabelToPbStageLabel(stageLabel)}, nil
}

// GetStageLabels gets all stage labels
func (s *StageLabelServer) GetStageLabels(ctx context.Context, req *pb.GetStageLabelsRequest) (*pb.GetStageLabelsResponse, error) {
	log.Println("GetStageLabels")

	stageLabels, err := s.StageLabelsStore.ReadStageLabels()
	if err != nil {
		return nil, err
	}

	return &pb.GetStageLabelsResponse{StageLabels: StageLabelToPbStageLabels(stageLabels)}, nil
}

// UpdateStageLabel updates a stage label
func (s *StageLabelServer) UpdateStageLabel(ctx context.Context, req *pb.UpdateStageLabelRequest) (*pb.UpdateStageLabelResponse, error) {
	log.Println("UpdateStageLabel")

	stageLabel := PbStageLabelToStageLabel(req.StageLabel)

	return &pb.UpdateStageLabelResponse{StageLabel: StageLabelToPbStageLabel(stageLabel)}, nil
}

// DeleteStageLabel deletes a stage label by id
func (s *StageLabelServer) DeleteStageLabel(ctx context.Context, req *pb.DeleteStageLabelRequest) (*pb.DeleteStageLabelResponse, error) {
	log.Println("DeleteStageLabel")

	err := s.StageLabelsStore.DeleteStageLabel(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteStageLabelResponse{
		Status: &pb.StatusP{
			Code:    200,
			Message: "Stage label deleted",
		},
	}, nil
}
