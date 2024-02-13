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

type PipelineStageServer struct {
	pb.UnimplementedPipelineStageServiceServer
	PipelineStageStore *DatabasePipelineStageStore
}

func NewPipelineStageServer(PipelineStageStore *DatabasePipelineStageStore) *PipelineStageServer {
	return &PipelineStageServer{
		PipelineStageStore: PipelineStageStore,
	}
}

// CreatePipelineStage creates a new pipeline stage
func (s *PipelineStageServer) CreatePipelineStage(ctx context.Context, req *pb.CreatePipelineStageRequest) (*pb.CreatePipelineStageResponse, error) {
	log.Println("CreatePipelineStage")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Pipeline Stage")
	}

	stage := PbStageToStage(req.Stage)
	stage.ID = uuid.New().String()
	stage.SerialNumber = utils.CreateSerialNumber("STG", stage.ID)
	stage.OrganisationID = contextData.OrganisationID
	stage.BranchId = contextData.BranchID
	stage.CreatedBy = contextData.UserID
	if stage.Owner == OwnerType(pb.OwnerType_ORGANIZATION) {
		stage.OwnerID = contextData.OrganisationID
	} else if stage.Owner == OwnerType(pb.OwnerType_BRANCH) {
		stage.OwnerID = contextData.BranchID
	} else if stage.Owner == OwnerType(pb.OwnerType_DEPARTMENT) {
		stage.OwnerID = stage.DepartmentID
	} else if stage.Owner == OwnerType(pb.OwnerType_TEAM) {
		stage.OwnerID = stage.TeamID
	} else if stage.Owner == OwnerType(pb.OwnerType_CREATOR) {
		stage.OwnerID = contextData.UserID
	}

	stage, err := s.PipelineStageStore.CreateStage(stage)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePipelineStageResponse{Stage: StageToPbStage(stage)}, nil
}

// GetPipelineStage gets a pipeline stage by id
func (s *PipelineStageServer) GetPipelineStage(ctx context.Context, req *pb.GetPipelineStageRequest) (*pb.GetPipelineStageResponse, error) {
	log.Println("GetPipelineStage")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Get Pipeline Stage")
	}

	stage, err := s.PipelineStageStore.ReadStage(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetPipelineStageResponse{Stage: StageToPbStage(stage)}, nil
}

// GetPipelineStages gets all pipeline stages for a pipeline
func (s *PipelineStageServer) GetPipelineStages(ctx context.Context, req *pb.GetPipelineStagesRequest) (*pb.GetPipelineStagesResponse, error) {
	log.Println("GetPipelineStages")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Get Pipeline Stages")
	}

	stages, err := s.PipelineStageStore.ReadStages(req.PipelineId)
	if err != nil {
		return nil, err
	}

	return &pb.GetPipelineStagesResponse{Stages: StagesToPbStages(convertReferences(stages))}, nil
}

// UpdatePipelineStage updates a pipeline stage
func (s *PipelineStageServer) UpdatePipelineStage(ctx context.Context, req *pb.UpdatePipelineStageRequest) (*pb.UpdatePipelineStageResponse, error) {
	log.Println("UpdatePipelineStage")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Update a Pipeline Stage")
	}

	stage := PbStageToStage(req.GetStage())
	stage.UpdatedBy = contextData.UserID
	stage, err := s.PipelineStageStore.UpdateStage(stage)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePipelineStageResponse{Stage: StageToPbStage(stage)}, nil
}

// DeletePipelineStage deletes a pipeline stage by id
func (s *PipelineStageServer) DeletePipelineStage(ctx context.Context, req *pb.DeletePipelineStageRequest) (*pb.DeletePipelineStageResponse, error) {
	log.Println("DeletePipelineStage")

	contextData, authErr := auth.GetRequestMetadata(ctx)
	if authErr != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Delete a Pipeline Stage")
	}
	err := s.PipelineStageStore.DeleteStage(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeletePipelineStageResponse{
		Status: &pb.Status{
			Code:    200,
			Message: "Pipeline stage deleted",
		},
	}, nil
}
