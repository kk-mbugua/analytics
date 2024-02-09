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
	DepartmentID   string
	BranchID       string         `gorm:"not null"`
	StageIDs       pq.StringArray `gorm:"not null"`
	OrganisationID string         `gorm:"not null"`
	CreatedBy      string         `gorm:"not null"`
	UpdatedBy      string         `gorm:"not null"`
	UpdatedAt      string         `gorm:"not null"`
	CreatedAt      string         `gorm:"not null"`
}
