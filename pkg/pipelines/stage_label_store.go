package pipelines

import "gorm.io/gorm"

type StageLabelStore interface {
	CreateStageLabel(stageLabel *StageLabel) (*StageLabel, error)
	ReadStageLabel(id string) (*StageLabel, error)
	ReadStageLabels(pipelineId string) ([]*StageLabel, error)
	UpdateStageLabel(stageLabel *StageLabel) (*StageLabel, error)
	DeleteStageLabel(id string) error
}

type DatabaseStageLabelStore struct {
	db *gorm.DB
}

func NewDatabaseStageLabelStore(db *gorm.DB) *DatabaseStageLabelStore {
	return &DatabaseStageLabelStore{db: db}
}

func (s *DatabaseStageLabelStore) CreateStageLabel(stageLabel *StageLabel) (*StageLabel, error) {
	result := s.db.Save(stageLabel).First(stageLabel)
	if result.Error != nil {
		return nil, result.Error
	}
	return stageLabel, nil
}

func (s *DatabaseStageLabelStore) ReadStageLabel(id string) (*StageLabel, error) {
	var stageLabel StageLabel
	err := s.db.Where("id = ?", id).First(&stageLabel).Error
	return &stageLabel, err
}

func (s *DatabaseStageLabelStore) ReadStageLabels(pipelineId string) ([]*StageLabel, error) {
	var stageLabels []*StageLabel
	err := s.db.Where("pipeline_id = ?", pipelineId).Find(&stageLabels).Error
	return stageLabels, err
}

func (s *DatabaseStageLabelStore) UpdateStageLabel(stageLabel *StageLabel) (*StageLabel, error) {
	result := s.db.Save(stageLabel).First(stageLabel)
	if result.Error != nil {
		return nil, result.Error
	}
	return stageLabel, nil
}

func (s *DatabaseStageLabelStore) DeleteStageLabel(id string) error {
	return s.db.Delete(&StageLabel{}, id).Error
}
