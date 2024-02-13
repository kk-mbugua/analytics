package pipelines

import (
	"main/pkg/proto/pb"

	"gorm.io/gorm"
)

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
		StageId:        lead.StageID,
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
		StageID:        lead.StageId,
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

func LeadOrganisation(orgId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orgId == "" {
			return db
		}
		return db.Where("organisation_id = ?", orgId)
	}
}

func LeadBranch(branchId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if branchId == "" {
			return db
		}
		return db.Where("branch_id = ?", branchId)
	}
}

func LeadOwner(ownerId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if ownerId == "" {
			return db
		}
		return db.Where("owner_id = ?", ownerId)
	}
}

func LeadDepartment(departmentId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if departmentId == "" {
			return db
		}
		return db.Where("department_id = ?", departmentId)
	}
}

func LeadTeam(teamId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if teamId == "" {
			return db
		}
		return db.Where("team_id = ?", teamId)
	}
}

func LeadPipeline(pipelineId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pipelineId == "" {
			return db
		}
		return db.Where("pipeline_ids @> ?", pipelineId)
	}
}

func LeadStage(stageId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if stageId == "" {
			return db
		}
		return db.Where("stage_id = ?", stageId)
	}
}

func LeadContact(contactId string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if contactId == "" {
			return db
		}
		return db.Where("contact_id = ?", contactId)
	}
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 || pageSize == 0 {
			return db
		}
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
