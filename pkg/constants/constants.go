package constants

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
