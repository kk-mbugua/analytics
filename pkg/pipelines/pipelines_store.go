package pipelines

import "gorm.io/gorm"

type PipelineStore interface {
	CreatePipeline(pipeline *Pipeline) (*Pipeline, error)
	ReadPipeline(id string) (*Pipeline, error)
	ReadPipelines(organisation_id string) ([]*Pipeline, error)
	UpdatePipeline(pipeline *Pipeline) (*Pipeline, error)
	DeletePipeline(id string) error
}

type DatabasePipelineStore struct {
	db *gorm.DB
}

func NewDatabasePipelineStore(db *gorm.DB) *DatabasePipelineStore {
	return &DatabasePipelineStore{db: db}
}

func (s *DatabasePipelineStore) CreatePipeline(pipeline *Pipeline) (*Pipeline, error) {
	result := s.db.Save(pipeline).First(pipeline)
	if result.Error != nil {
		return nil, result.Error
	}
	return pipeline, nil
}

func (s *DatabasePipelineStore) ReadPipeline(id string) (*Pipeline, error) {
	var pipeline Pipeline
	err := s.db.Where("id = ?", id).Preload("Stages").First(&pipeline).Error
	return &pipeline, err
}

func (s *DatabasePipelineStore) ReadPipelines(organisation_id string) ([]*Pipeline, error) {
	var pipelines []*Pipeline
	err := s.db.Where("organisation_id = ?", organisation_id).Find(&pipelines).Error
	return pipelines, err
}

func (s *DatabasePipelineStore) UpdatePipeline(pipeline *Pipeline) (*Pipeline, error) {
	result := s.db.Save(pipeline).First(pipeline)
	if result.Error != nil {
		return nil, result.Error
	}
	return pipeline, nil
}

func (s *DatabasePipelineStore) DeletePipeline(id string) error {
	return s.db.Delete(&Pipeline{}, id).Error
}
