package pipelines

import (
	"gorm.io/gorm"
)

type OwnerType int

const (
	UNKNOWN      OwnerType = iota
	TEAM                   = "TEAM"
	CREATOR                = "CREATOR"
	BRANCH                 = "BRANCH"
	DEPARTMENT             = "DEPARTMENT"
	ORGANISATION           = "ORGANISATION"
)

type PipelineTypes int

const (
	UnknownPipelineType PipelineTypes = iota
	SALES                             = "SALES"
	MARKETING                         = "MARKETING"
	CUSTOMER_SUPPORT                  = "CUSTOMER_SUPPORT"
	NON_PROFIT                        = "NON_PROFIT"
	CUSTOM                            = "CUSTOM"
)

type Pipeline struct {
	gorm.Model
	ID                    string `gorm:"primary_key;type:uuid"`
	Name                  string `gorm:"not null"`
	Description           string `gorm:"not null"`
	PipelineStages        []Stage
	PipelineType          PipelineTypes `gorm:"not null"`
	CustomTypeName        string
	CustomTypeDescription string
	OrganisationID        string    `gorm:"not null"`
	SerialNumber          string    `gorm:"not null"`
	BranchID              string    `gorm:"not null"`
	OwnerID               string    `gorm:"not null"`
	Owner                 OwnerType `gorm:"not null"`
	CreatedBy             string    `gorm:"not null"`
	UpdatedBy             string    `gorm:"not null"`
	UpdatedAt             string    `gorm:"not null"`
	CreatedAt             string    `gorm:"not null"`
	BranchId              string    `gorm:"not null"`
	DepartmentID          string
	TeamID                string
}

type Stage struct {
	gorm.Model
	ID             string `gorm:"primary_key;type:uuid"`
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	PipelineID     string `gorm:"not null;foreignkey:PipelineID"`
	Index          int32  `gorm:"not null"`
	StageLabelID   string `gorm:"not null;foreignkey:StageLabelID"`
	StageLabel     StageLabel
	OrganisationID string `gorm:"not null"`
	SerialNumber   string `gorm:"not null"`
	CreatedBy      string `gorm:"not null"`
	UpdatedBy      string `gorm:"not null"`
	UpdatedAt      string `gorm:"not null"`
	CreatedAt      string `gorm:"not null"`
	BranchId       string `gorm:"not null"`
	DepartmentId   string
	TeamID         string
	OwnerID        string
	Owner          OwnerType
	DepartmentID   string
	CustomFields   []CustomField
}

type StageLabel struct {
	gorm.Model
	ID             string `gorm:"primary_key;type:uuid"`
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	Color          string `gorm:"not null"`
	Banner         string `gorm:"not null"`
	OrganisationID string `gorm:"not null"`
	SerialNumber   string `gorm:"not null"`
	CreatedBy      string `gorm:"not null"`
	UpdatedBy      string `gorm:"not null"`
	UpdatedAt      string `gorm:"not null"`
	CreatedAt      string `gorm:"not null"`
	BranchID       string `gorm:"not null"`
	DepartmentId   string
}

type CustomField struct {
	gorm.Model
	ID             string `gorm:"primary_key;type:uuid"`
	Name           string
	Type           string
	StageID        string `gorm:"not null;foreignkey:StageID"`
	OrganisationID string `gorm:"not null"`
}
