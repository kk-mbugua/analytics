package pipelines

import (
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type LeadStore interface {
	CreateLead(lead *Lead) (*Lead, error)
	GetLead(id string) (*Lead, error)
	GetLeads(page int, pageSize int, filters *LeadFilters) ([]*Lead, error)
	UpdateLead(lead *Lead) (*Lead, error)
	MoveLeadToStage(leadID, stageID string) error
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

func (d *DatabaseLeadStore) GetLeads(page int, pageSize int, filters *LeadFilters) ([]*Lead, error) {
	var leads []*Lead
	result := d.db.Scopes(
		LeadBranch(filters.BranchID),
		LeadOrganisation(filters.OrganisationID),
		LeadOwner(filters.OwnerID),
		LeadDepartment(filters.DepartmentID),
		LeadTeam(filters.TeamID),
		LeadPipeline(filters.PipelineID),
		LeadStage(filters.StageID),
		LeadContact(filters.ContactID),
		Paginate(page, pageSize),
	).Find(&leads, filters)
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
func (d *DatabaseLeadStore) MoveLeadToStage(leadID, stageID string) error {
	// Get the stage from the database
	stage := &Stage{}
	if err := d.db.First(stage, stageID).Error; err != nil {
		return err
	}

	// Get the lead from the database
	lead := &Lead{}
	if err := d.db.First(lead, leadID).Error; err != nil {
		return err
	}

	// Update the lead's stage
	lead.StageID = stageID

	// Add the custom fields for the new stage
	for _, field := range stage.CustomFields {
		lead.CustomFields = append(lead.CustomFields, LeadCustomFields{
			LeadID:        leadID,
			CustomFieldID: field.ID,
			Value:         "",
		})
	}

	// Save the lead back to the database
	if err := d.db.Save(lead).Error; err != nil {
		return err
	}

	return nil
}

func (d *DatabaseLeadStore) UpdateLeadCustomFieldValue(leadID, customFieldID, value string) error {
	// Get the lead from the database
	lead := &Lead{}
	if err := d.db.First(lead, leadID).Error; err != nil {
		return err
	}

	// Get the custom field from the database
	field := &CustomField{}
	if err := d.db.First(field, customFieldID).Error; err != nil {
		return err
	}

	// Update the lead's custom field value
	for i, f := range lead.CustomFields {
		if f.CustomFieldID == customFieldID {
			lead.CustomFields[i].Value = value
		}
	}

	// Save the lead back to the database
	if err := d.db.Save(lead).Error; err != nil {
		return err
	}

	return nil
}
