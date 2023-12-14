package pipelines

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type StageType int

const (
	UnknownStageType StageType = iota
	Lead                       = "LEAD"
	Account                    = "ACCOUNT"
)

type Pipeline struct {
	gorm.Model
	ID          string `gorm:"primary_key;type:uuid"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
}

type Stage struct {
	gorm.Model
	ID          string    `gorm:"primary_key;type:uuid"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	PipelineID  string    `gorm:"not null;foreignkey:PipelineID"`
	Index       int32     `gorm:"not null"`
	Type        StageType `gorm:"not null"`
}

type RuleVariableType int

const (
	UnknownRuleType RuleVariableType = iota
	Text                             = "TEXT"
	Date                             = "DATE"
	Number                           = "NUMBER"
	Bool                             = "BOOL"
)

type OperandType int

const (
	UnknownOperandType OperandType = iota
	IsEquals                       = "ISEQUALS"
	Between                        = "BETWEEN"
	NotBetween                     = "NOTBETWEEN"
	GreaterThan                    = "GREATERTHAN"
	LessThan                       = "LESSTHAN"
	In                             = "IN"
	NotIn                          = "NOTIN"
	Contains                       = "CONTAINS"
	NotContains                    = "NOTCONTAINS"
	StartsWith                     = "STARTSWITH"
	NotStartsWith                  = "NOTSTARTSWITH"
	EndsWith                       = "ENDSWITH"
	NotEndsWith                    = "NOTENDSWITH"
)

type Rule struct {
	gorm.Model
	ID           string           `gorm:"primary_key;type:uuid"`
	Variable     string           `gorm:"not null"`
	VariableType RuleVariableType `gorm:"not null"`
	Operand      OperandType      `gorm:"not null"`
	Value        pq.StringArray   `gorm:"not null"`
	StageID      string           `gorm:"not null;foreignkey:StageID"`
	Index        int32            `gorm:"not null"`
}
