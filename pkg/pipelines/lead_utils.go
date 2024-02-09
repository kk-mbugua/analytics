package pipelines

import "main/pkg/proto/pb"

func LeadToPbLead(lead *Lead) *pb.Lead {
	return &pb.Lead{
		Id:             lead.ID,
		SerialNumber:   lead.SerialNumber,
		ContactId:      lead.ContactID,
		PipelineIds:    lead.PipelineIDs,
		Owner:          pb.OwnerType(lead.Owner),
		OwnerId:        lead.OwnerID,
		DepartmentId:   lead.DepartmentID,
		BranchId:       lead.BranchID,
		StageIds:       lead.StageIDs,
		OrganisationId: lead.OrganisationID,
		CreatedBy:      lead.CreatedBy,
		UpdatedBy:      lead.UpdatedBy,
		UpdatedAt:      lead.UpdatedAt,
		CreatedAt:      lead.CreatedAt,
	}
}

func PbLeadToLead(lead *pb.Lead) *Lead {
	return &Lead{
		ID:             lead.Id,
		SerialNumber:   lead.SerialNumber,
		ContactID:      lead.ContactId,
		PipelineIDs:    lead.PipelineIds,
		Owner:          OwnerType(lead.Owner),
		OwnerID:        lead.OwnerId,
		DepartmentID:   lead.DepartmentId,
		BranchID:       lead.BranchId,
		StageIDs:       lead.StageIds,
		OrganisationID: lead.OrganisationId,
		CreatedBy:      lead.CreatedBy,
		UpdatedBy:      lead.UpdatedBy,
		UpdatedAt:      lead.UpdatedAt,
		CreatedAt:      lead.CreatedAt,
	}
}

func PbLeadsToLeads(leads []*pb.Lead) []*Lead {
	var result []*Lead
	for _, lead := range leads {
		result = append(result, PbLeadToLead(lead))
	}
	return result
}

func LeadsToPbLeads(leads []*Lead) []*pb.Lead {
	var result []*pb.Lead
	for _, lead := range leads {
		result = append(result, LeadToPbLead(lead))
	}
	return result
}
