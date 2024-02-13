package constants

import (
	"time"
)

type StageType int

const (
	UnknownStageType StageType = iota
	LEAD                       = "LEAD"
	ACCOUNT                    = "ACCOUNT"
	OPPORTUNITY                = "OPPORTUNITY"
	CUSTOMER                   = "CUSTOMER"
	LOST                       = "LOST"
	CUSTOMLABEL                = "CUSTOMLABEL"
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

type Consts struct {
	FALSE                      string
	TRUE                       string
	FALSE_BOOL                 bool
	TRUE_BOOL                  bool
	ContactsRequiredFields     []string
	OrganisationRequiredFields []string
	TagsRequiredFields         []string
	LabelsRequiredFields       []string
}

func NewConsts() *Consts {
	return &Consts{
		FALSE:                      "false",
		TRUE:                       "true",
		FALSE_BOOL:                 false,
		TRUE_BOOL:                  true,
		ContactsRequiredFields:     []string{"FirstName", "LastName"},
		OrganisationRequiredFields: []string{"name", "description", "first_name", "last_name", "username", "email", "password", "phone"},
		TagsRequiredFields:         []string{"Name"},
		LabelsRequiredFields:       []string{"Name"},
	}
}

func TIME_NOW() time.Time {
	return time.Now()
}

func (c *Consts) BoolToString(b bool) string {
	if b {
		return c.TRUE
	}
	return c.FALSE
}
