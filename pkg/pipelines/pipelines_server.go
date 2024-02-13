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
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Pipeline")
	}

	pipeline := PbPipelineToPipeline(req.Pipeline)
	pipeline.ID = uuid.New().String()
	pipeline.SerialNumber = utils.CreateSerialNumber("PIP", pipeline.ID)
	pipeline.OrganisationID = contextData.OrganisationID
	pipeline.CreatedBy = contextData.UserID
	pipeline.BranchID = contextData.BranchID
	if pipeline.Owner == OwnerType(pb.OwnerType_ORGANIZATION) {
		pipeline.OwnerID = contextData.OrganisationID
	} else if pipeline.Owner == OwnerType(pb.OwnerType_BRANCH) {
		pipeline.OwnerID = contextData.BranchID
	} else if pipeline.Owner == OwnerType(pb.OwnerType_DEPARTMENT) {
		pipeline.OwnerID = pipeline.DepartmentID
	} else if pipeline.Owner == OwnerType(pb.OwnerType_TEAM) {
		pipeline.OwnerID = pipeline.TeamID
	} else if pipeline.Owner == OwnerType(pb.OwnerType_CREATOR) {
		pipeline.OwnerID = contextData.UserID
	}
	pipeline, err = s.PipelineStore.CreatePipeline(pipeline)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePipelineResponse{
		Pipeline: PipelineToPbPipeline(pipeline),
		Status: &pb.Status{
			Code:    200,
			Message: "Pipeline created Successfully",
		},
	}, nil
}

func (s *PipelineServer) GetPipeline(ctx context.Context, req *pb.GetPipelineRequest) (*pb.GetPipelineResponse, error) {
	log.Println("GetPipeline")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Get a Pipeline")
	}

	pipeline, err := s.PipelineStore.ReadPipeline(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetPipelineResponse{
		Pipeline: PipelineToPbPipeline(pipeline),
		Status: &pb.Status{
			Code:    200,
			Message: "Pipeline retrieved Successfully",
		},
	}, nil
}

func (s *PipelineServer) GetPipelines(ctx context.Context, req *pb.GetPipelinesRequest) (*pb.GetPipelinesResponse, error) {
	log.Println("GetPipelines")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to List Pipelines")
	}

	pipelines, err := s.PipelineStore.ReadPipelines(req.GetOrganisationId())
	if err != nil {
		return nil, err
	}

	return &pb.GetPipelinesResponse{
		Pipelines: PipelinesToPbPipelines(pipelines),
		Status: &pb.Status{
			Code:    200,
			Message: "Pipelines retrieved Successfully",
		},
	}, nil
}

func (s *PipelineServer) UpdatePipeline(ctx context.Context, req *pb.UpdatePipelineRequest) (*pb.UpdatePipelineResponse, error) {
	log.Println("UpdatePipeline")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Update a Pipeline")
	}

	pipeline := PbPipelineToPipeline(req.Pipeline)
	pipeline.UpdatedBy = contextData.UserID
	pipeline, err = s.PipelineStore.UpdatePipeline(pipeline)
	if err != nil {
		return nil, err
	}

	return &pb.UpdatePipelineResponse{
		Pipeline: PipelineToPbPipeline(pipeline),
		Status: &pb.Status{
			Code:    200,
			Message: "Pipeline updated Successfully",
		},
	}, nil
}

func (s *PipelineServer) DeletePipeline(ctx context.Context, req *pb.DeletePipelineRequest) (*pb.DeletePipelineResponse, error) {
	log.Println("DeletePipeline")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Delete a Pipeline")
	}

	if err := s.PipelineStore.DeletePipeline(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.DeletePipelineResponse{
		Status: &pb.Status{
			Code:    200,
			Message: "Pipeline deleted",
		},
	}, nil
}
