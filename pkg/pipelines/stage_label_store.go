package pipelines

import "gorm.io/gorm"

type StageLabelStore interface {
	CreateStageLabel(stageLabel *StageLabel) error
	ReadStageLabel(id string) (*StageLabel, error)
	ReadStageLabels() ([]*StageLabel, error)
	UpdateStageLabel(stageLabel *StageLabel) error
	DeleteStageLabel(id string) error
}

type DatabaseStageLabelStore struct {
	db *gorm.DB
}

func NewDatabaseStageLabelStore(db *gorm.DB) *DatabaseStageLabelStore {
	return &DatabaseStageLabelStore{db: db}
}

func (s *DatabaseStageLabelStore) CreateStageLabel(stageLabel *StageLabel) error {
	return s.db.Create(stageLabel).Error
}

func (s *DatabaseStageLabelStore) ReadStageLabel(id string) (*StageLabel, error) {
	var stageLabel StageLabel
	err := s.db.Where("id = ?", id).First(&stageLabel).Error
	return &stageLabel, err
}

func (s *DatabaseStageLabelStore) ReadStageLabels() ([]*StageLabel, error) {
	var stageLabels []*StageLabel
	err := s.db.Find(&stageLabels).Error
	return stageLabels, err
}

func (s *DatabaseStageLabelStore) UpdateStageLabel(stageLabel *StageLabel) error {
	return s.db.Save(stageLabel).Error
}

func (s *DatabaseStageLabelStore) DeleteStageLabel(id string) error {
	return s.db.Delete(&StageLabel{}, id).Error
}
