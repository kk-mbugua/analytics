package pipelines

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	ID             string         `gorm:"primaryKey"`
	SerialNumber   string         `gorm:"not null"`
	ContactID      string         `gorm:"not null"`
	PipelineIDs    pq.StringArray `gorm:"type:text[]"`
	OwnerID        string         `gorm:"not null"`
	Owner          OwnerType      `gorm:"not null"`
	CustomFields   []LeadCustomFields
	TeamID         string
	DepartmentID   string
	BranchID       string `gorm:"not null"`
	StageID        string `gorm:"not null"`
	OrganisationID string `gorm:"not null"`
	CreatedBy      string `gorm:"not null"`
	UpdatedBy      string `gorm:"not null"`
	UpdatedAt      string `gorm:"not null"`
	CreatedAt      string `gorm:"not null"`
}

type LeadCustomFields struct {
	gorm.Model
	LeadID      string `gorm:"not null"`
	FieldID     string `gorm:"not null"`
	CustomField CustomField
	Value       string `gorm:"not null"`
}

type LeadFilters struct {
	BranchID       string
	OrganisationID string
	OwnerID        string
	StageID        string
	PipelineID     string
	DepartmentID   string
	TeamID         string
	ContactID      string
}

type LeadQualifiers struct {
	gorm.Model
	QualifiersName string `gorm:"not null"`
	QualifiersType string `gorm:"not null"`
	Value          string `gorm:"not null;default:false"`
	StageID        string `gorm:"not null"`
	LeadID         string `gorm:"not null"`
}
