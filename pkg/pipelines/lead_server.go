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

// TODO: Refactor methods

type LeadServer struct {
	pb.UnimplementedLeadServiceServer
	LeadStore LeadStore
}

func NewLeadServer(leadStore LeadStore) *LeadServer {
	return &LeadServer{LeadStore: leadStore}
}

func (l *LeadServer) CreateLead(ctx context.Context, req *pb.CreateLeadRequest) (*pb.CreateLeadResponse, error) {
	log.Println("CreateLead")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Lead")
	}
	log.Println("contextData", contextData)

	lead := PbLeadToLead(req.GetLead())
	lead.ID = uuid.New().String()
	lead.SerialNumber = utils.CreateSerialNumber("LEA", lead.ID)
	lead.OrganisationID = contextData.OrganisationID
	lead.CreatedBy = contextData.UserID
	lead.BranchID = contextData.BranchID
	if lead.Owner == OwnerType(pb.OwnerType_ORGANIZATION) {
		lead.OwnerID = contextData.OrganisationID
	} else if lead.Owner == OwnerType(pb.OwnerType_BRANCH) {
		lead.OwnerID = contextData.BranchID
	} else if lead.Owner == OwnerType(pb.OwnerType_DEPARTMENT) {
		lead.OwnerID = lead.DepartmentID
	} else if lead.Owner == OwnerType(pb.OwnerType_TEAM) {
		lead.OwnerID = lead.TeamID
	} else if lead.Owner == OwnerType(pb.OwnerType_CREATOR) {
		lead.OwnerID = contextData.UserID
	}
	createdLead, err := l.LeadStore.CreateLead(lead)
	if err != nil {
		return nil, err
	}

	return &pb.CreateLeadResponse{
		Lead: LeadToPbLead(createdLead),
		Status: &pb.Status{
			Code:    200,
			Message: "Lead created Successfully",
		},
	}, nil
}

func (l *LeadServer) GetLead(ctx context.Context, req *pb.GetLeadRequest) (*pb.GetLeadResponse, error) {
	log.Println("GetLead")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Lead")
	}

	lead, err := l.LeadStore.GetLead(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetLeadResponse{
		Lead: LeadToPbLead(lead),
		Status: &pb.Status{
			Code:    200,
			Message: "Lead Fetched Successfully",
		},
	}, nil
}

func (l *LeadServer) GetLeads(ctx context.Context, req *pb.ListLeadRequest) (*pb.ListLeadResponse, error) {
	log.Println("GetLeadsByOwnerID")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Lead")
	}

	leadFilters := &LeadFilters{
		BranchID:       req.GetBranchId(),
		OrganisationID: contextData.OrganisationID,
		OwnerID:        req.GetOwnerId(),
		StageID:        req.GetStageId(),
		PipelineID:     req.GetPipelineId(),
		DepartmentID:   req.GetDepartmentId(),
		TeamID:         req.GetTeamId(),
		ContactID:      req.GetContactId(),
	}

	leads, err := l.LeadStore.GetLeads(int(req.GetPage()), int(req.GetPageSize()), leadFilters)
	if err != nil {
		return nil, err
	}

	return &pb.ListLeadResponse{
		Leads: LeadsToPbLeads(leads),
		Status: &pb.Status{
			Code:    200,
			Message: "Leads Fetched Successfully",
		},
	}, nil
}

func (l *LeadServer) UpdateLead(ctx context.Context, req *pb.UpdateLeadRequest) (*pb.UpdateLeadResponse, error) {
	log.Println("UpdateLead")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Lead")
	}

	lead := PbLeadToLead(req.GetLead())
	lead.UpdatedBy = contextData.UserID
	updatedLead, err := l.LeadStore.UpdateLead(lead)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateLeadResponse{
		Lead: LeadToPbLead(updatedLead),
		Status: &pb.Status{
			Code:    200,
			Message: "Lead updated Successfully",
		},
	}, nil
}

func (l *LeadServer) DeleteLead(ctx context.Context, req *pb.DeleteLeadRequest) (*pb.DeleteLeadResponse, error) {
	log.Println("DeleteLead")
	contextData, err := auth.GetRequestMetadata(ctx)
	if err != nil || contextData.RequestAuth == constants.NewConsts().FALSE {
		return nil, status.Errorf(403, "Forbidden, You do not have permission to Create a Lead")
	}

	if err := l.LeadStore.DeleteLead(req.GetId()); err != nil {
		return nil, err
	}

	return &pb.DeleteLeadResponse{
		Status: &pb.Status{
			Code:    200,
			Message: "Lead deleted",
		},
	}, nil
}
