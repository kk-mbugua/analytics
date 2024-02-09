package pipelines

import (
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type LeadStore interface {
	CreateLead(lead *Lead) (*Lead, error)
	GetLead(id string) (*Lead, error)
	GetLeadsByOwnerID(ownerID string) ([]*Lead, error)
	GetLeadsByBranchID(branchID string) ([]*Lead, error)
	GetLeadsByDepartmentID(departmentID string) ([]*Lead, error)
	GetLeadsByOrganisationID(organisationID string) ([]*Lead, error)
	GetLeadsByPipelineID(pipelineID string) ([]*Lead, error)
	GetLeadsByStageID(stageID string) ([]*Lead, error)
	UpdateLead(lead *Lead) (*Lead, error)
	DeleteLead(id string) error
}

type DatabaseLeadStore struct {
	db *gorm.DB
}

func NewDatabaseLeadStore(db *gorm.DB) *DatabaseLeadStore {
	return &DatabaseLeadStore{db: db}
}

func (d *DatabaseLeadStore) CreateLead(lead *Lead) (*Lead, error) {
	result := d.db.Create(lead).First(lead)
	if result.Error != nil {
		return nil, status.Error(500, "Failed to create lead")
	}
	return lead, nil
}

func (d *DatabaseLeadStore) GetLead(id string) (*Lead, error) {
	var lead *Lead
	result := d.db.First(&lead, "id = ?", id)
	if result.Error != nil {
		return nil, status.Error(404, "Lead not found")
	}
	return lead, nil
}

func (d *DatabaseLeadStore) GetLeadsByOwnerID(ownerID string) ([]*Lead, error) {
	var leads []*Lead
	result := d.db.Find(&leads, "owner_id = ?", ownerID)
	if result.Error != nil {
		return nil, status.Error(404, "Leads not found")
	}
	return leads, nil
}

func (d *DatabaseLeadStore) GetLeadsByBranchID(branchID string) ([]*Lead, error) {
	var leads []*Lead
	result := d.db.Find(&leads, "branch_id = ?", branchID)
	if result.Error != nil {
		return nil, status.Error(404, "Leads not found")
	}
	return leads, nil
}

func (d *DatabaseLeadStore) GetLeadsByDepartmentID(departmentID string) ([]*Lead, error) {
	var leads []*Lead
	result := d.db.Find(&leads, "department_id = ?", departmentID)
	if result.Error != nil {
		return nil, status.Error(404, "Leads not found")
	}
	return leads, nil
}

func (d *DatabaseLeadStore) GetLeadsByOrganisationID(organisationID string) ([]*Lead, error) {
	var leads []*Lead
	result := d.db.Find(&leads, "organisation_id = ?", organisationID)
	if result.Error != nil {
		return nil, status.Error(404, "Leads not found")
	}
	return leads, nil
}

func (d *DatabaseLeadStore) GetLeadsByPipelineID(pipelineID string) ([]*Lead, error) {
	var leads []*Lead
	result := d.db.Find(&leads, "pipeline_ids = ?", pipelineID)
	if result.Error != nil {
		return nil, status.Error(404, "Leads not found")
	}
	return leads, nil
}

func (d *DatabaseLeadStore) GetLeadsByStageID(stageID string) ([]*Lead, error) {

	var leads []*Lead
	result := d.db.Find(&leads, "stage_ids = ?", stageID)
	if result.Error != nil {
		return nil, status.Error(404, "Leads not found")
	}
	return leads, nil
}

func (d *DatabaseLeadStore) UpdateLead(lead *Lead) (*Lead, error) {
	result := d.db.Save(lead)
	if result.Error != nil {
		return nil, status.Error(500, "Failed to update lead")
	}
	return lead, nil
}

func (d *DatabaseLeadStore) DeleteLead(id string) error {
	result := d.db.Delete(&Lead{}, "id = ?", id)
	if result.Error != nil {
		return status.Error(500, "Failed to delete lead")
	}
	return nil
}
