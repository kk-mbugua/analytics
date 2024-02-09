package pipelines

import (
	"gorm.io/gorm"
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
	ID             string `gorm:"primary_key;type:uuid"`
	Name           string `gorm:"not null"`
	Description    string `gorm:"not null"`
	PipelineStages []Stage
	PipelineType   PipelineTypes `gorm:"not null"`
	CustomTypeName string
	OrganisationID string `gorm:"not null"`
	SerialNumber   string `gorm:"not null"`
	CreatedBy      string `gorm:"not null"`
	UpdatedBy      string `gorm:"not null"`
	UpdatedAt      string `gorm:"not null"`
	CreatedAt      string `gorm:"not null"`
	BranchId       string `gorm:"not null"`
	DepartmentId   string
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
	BranchId       string `gorm:"not null"`
	DepartmentId   string
}
