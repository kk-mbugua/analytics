package pipelines

import "gorm.io/gorm"

type PipelineStageStore interface {
	CreateStage(stage *Stage) error
	ReadStage(id string) (*Stage, error)
	ReadStages(pipeline_id string) ([]*Stage, error)
	UpdateStage(stage *Stage) error
	DeleteStage(id string) error
}

type DatabasePipelineStageStore struct {
	db *gorm.DB
}

func NewDatabasePipelineStageStore(db *gorm.DB) *DatabasePipelineStageStore {
	return &DatabasePipelineStageStore{db: db}
}

func (s *DatabasePipelineStageStore) CreateStage(stage *Stage) error {
	return s.db.Create(stage).Error
}

func (s *DatabasePipelineStageStore) ReadStage(id string) (*Stage, error) {
	var stage Stage
	err := s.db.Where("id = ?", id).First(&stage).Error
	return &stage, err
}

func (s *DatabasePipelineStageStore) ReadStages(pipeline_id string) ([]*Stage, error) {
	var stages []*Stage
	err := s.db.Where("pipeline_id = ?", pipeline_id).Find(&stages).Error
	return stages, err
}

func (s *DatabasePipelineStageStore) UpdateStage(stage *Stage) error {
	return s.db.Save(stage).Error
}

func (s *DatabasePipelineStageStore) DeleteStage(id string) error {
	return s.db.Delete(&Stage{}, id).Error
}
