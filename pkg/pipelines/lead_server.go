package pipelines

import "log"

// TODO: Refactor methods

type LeadServer struct {
	LeadStore LeadStore
}

func NewLeadServer(leadStore LeadStore) *LeadServer {
	return &LeadServer{LeadStore: leadStore}
}

func (l *LeadServer) CreateLead(lead *Lead) (*Lead, error) {
	log.Println("CreateLead")

	createdLead, err := l.LeadStore.CreateLead(lead)
	if err != nil {
		return nil, err
	}

	return createdLead, nil
}

func (l *LeadServer) GetLead(id string) (*Lead, error) {
	log.Println("GetLead")

	lead, err := l.LeadStore.GetLead(id)
	if err != nil {
		return nil, err
	}

	return lead, nil
}

func (l *LeadServer) GetLeadsByOwnerID(ownerID string) ([]*Lead, error) {
	log.Println("GetLeadsByOwnerID")

	leads, err := l.LeadStore.GetLeadsByOwnerID(ownerID)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (l *LeadServer) GetLeadsByBranchID(branchID string) ([]*Lead, error) {
	log.Println("GetLeadsByBranchID")

	leads, err := l.LeadStore.GetLeadsByBranchID(branchID)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (l *LeadServer) GetLeadsByDepartmentID(departmentID string) ([]*Lead, error) {
	log.Println("GetLeadsByDepartmentID")

	leads, err := l.LeadStore.GetLeadsByDepartmentID(departmentID)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (l *LeadServer) GetLeadsByOrganisationID(organisationID string) ([]*Lead, error) {
	log.Println("GetLeadsByOrganisationID")

	leads, err := l.LeadStore.GetLeadsByOrganisationID(organisationID)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (l *LeadServer) GetLeadsByPipelineID(pipelineID string) ([]*Lead, error) {
	log.Println("GetLeadsByPipelineID")

	leads, err := l.LeadStore.GetLeadsByPipelineID(pipelineID)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (l *LeadServer) GetLeadsByStageID(stageID string) ([]*Lead, error) {
	log.Println("GetLeadsByStageID")

	leads, err := l.LeadStore.GetLeadsByStageID(stageID)
	if err != nil {
		return nil, err
	}

	return leads, nil
}

func (l *LeadServer) UpdateLead(lead *Lead) (*Lead, error) {
	log.Println("UpdateLead")

	updatedLead, err := l.LeadStore.UpdateLead(lead)
	if err != nil {
		return nil, err
	}

	return updatedLead, nil
}

func (l *LeadServer) DeleteLead(id string) error {
	log.Println("DeleteLead")

	err := l.LeadStore.DeleteLead(id)
	if err != nil {
		return err
	}

	return nil
}
