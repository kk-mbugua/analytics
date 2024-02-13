package pipelines

import (
	"context"
	"log"
	utils "main/pkg/Util"
	"main/pkg/auth"
	"main/pkg/constants"
	"main/pkg/proto/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/status"
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

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Stage Label")
	}
	stageLabel := PbStageLabelToStageLabel(req.GetStageLabel())
	stageLabel.ID = uuid.New().String()
	stageLabel.SerialNumber = utils.CreateSerialNumber("STGL", stageLabel.ID)
	stageLabel.OrganisationID = contextData.OrganisationID
	stageLabel.BranchID = contextData.BranchID
	stageLabel.CreatedBy = contextData.UserID

	newStageLabel, err := s.StageLabelsStore.CreateStageLabel(stageLabel)
	if err != nil {
		return nil, err
	}
	return &pb.CreateStageLabelResponse{StageLabel: StageLabelToPbStageLabel(newStageLabel)}, nil
}

// GetStageLabel gets a stage label by id
func (s *StageLabelServer) GetStageLabel(ctx context.Context, req *pb.GetStageLabelRequest) (*pb.GetStageLabelResponse, error) {
	log.Println("GetStageLabel")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to access this resource")
	}

	stageLabel, err := s.StageLabelsStore.ReadStageLabel(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetStageLabelResponse{StageLabel: StageLabelToPbStageLabel(stageLabel)}, nil
}

// GetStageLabels gets all stage labels
func (s *StageLabelServer) GetStageLabels(ctx context.Context, req *pb.GetStageLabelsRequest) (*pb.GetStageLabelsResponse, error) {
	log.Println("GetStageLabels")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to access this resource")
	}

	stageLabels, err := s.StageLabelsStore.ReadStageLabels(req.GetPipelineId())
	if err != nil {
		return nil, err
	}

	return &pb.GetStageLabelsResponse{StageLabels: StageLabelToPbStageLabels(stageLabels)}, nil
}

// UpdateStageLabel updates a stage label
func (s *StageLabelServer) UpdateStageLabel(ctx context.Context, req *pb.UpdateStageLabelRequest) (*pb.UpdateStageLabelResponse, error) {
	log.Println("UpdateStageLabel")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to access this resource")
	}

	stageLabel := PbStageLabelToStageLabel(req.GetStageLabel())
	stageLabel.UpdatedBy = contextData.UserID

	updatedStageLabel, err := s.StageLabelsStore.UpdateStageLabel(stageLabel)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateStageLabelResponse{StageLabel: StageLabelToPbStageLabel(updatedStageLabel)}, nil
}

// DeleteStageLabel deletes a stage label by id
func (s *StageLabelServer) DeleteStageLabel(ctx context.Context, req *pb.DeleteStageLabelRequest) (*pb.DeleteStageLabelResponse, error) {
	log.Println("DeleteStageLabel")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to access this resource")
	}

	if err := s.StageLabelsStore.DeleteStageLabel(req.Id); err != nil {
		return nil, err
	}

	return &pb.DeleteStageLabelResponse{
		Status: &pb.Status{
			Code:    200,
			Message: "Stage label deleted",
		},
	}, nil
}
