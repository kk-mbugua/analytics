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

type Consts struct {
	FALSE      string
	TRUE       string
	FALSE_BOOL bool
	TRUE_BOOL  bool
}

func NewConsts() *Consts {
	return &Consts{
		FALSE:      "false",
		TRUE:       "true",
		FALSE_BOOL: false,
		TRUE_BOOL:  true,
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
