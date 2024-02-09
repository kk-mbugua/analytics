package pipelines

import "main/pkg/proto/pb"

func PipelineToPbPipeline(pipeline *Pipeline) *pb.Pipeline {
	return &pb.Pipeline{
		Id:             pipeline.ID,
		Name:           pipeline.Name,
		Stages:         StagesToPbStages(convertToReferences(pipeline.PipelineStages)),
		Description:    pipeline.Description,
		PipelineType:   pb.Pipeline_PipelineType(pipeline.PipelineType),
		OrganisationId: pipeline.OrganisationID,
	}
}

func PbPipelineToPipeline(pipeline *pb.Pipeline) *Pipeline {
	return &Pipeline{
		ID:             pipeline.Id,
		Name:           pipeline.Name,
		Description:    pipeline.Description,
		PipelineStages: PbStagesToStages(pipeline.Stages),
		PipelineType:   PipelineTypes(pipeline.GetPipelineType()),
		OrganisationID: pipeline.OrganisationId,
	}
}

func StageToPbStage(stage *Stage) *pb.PipelineStage {
	return &pb.PipelineStage{
		Id:             stage.ID,
		Name:           stage.Name,
		Description:    stage.Description,
		PipelineId:     stage.PipelineID,
		Index:          stage.Index,
		StageLabelId:   stage.StageLabelID,
		StageLabel:     StageLabelToPbStageLabel(&stage.StageLabel),
		SerialNumber:   stage.SerialNumber,
		OrganisationId: stage.OrganisationID,
		CreatedBy:      stage.CreatedBy,
		UpdatedBy:      stage.UpdatedBy,
		UpdatedAt:      stage.UpdatedAt,
		CreatedAt:      stage.CreatedAt,
		BranchId:       stage.BranchId,
		DepartmentId:   stage.DepartmentId,
	}
}

func StageLabelToPbStageLabel(stageLabel *StageLabel) *pb.StageLabel {
	return &pb.StageLabel{
		Id:          stageLabel.ID,
		Name:        stageLabel.Name,
		Description: stageLabel.Description,
		Color:       stageLabel.Color,
		Banner:      stageLabel.Banner,
	}
}

func PbStageLabelToStageLabel(stageLabel *pb.StageLabel) *StageLabel {
	return &StageLabel{
		ID:          stageLabel.Id,
		Name:        stageLabel.Name,
		Description: stageLabel.Description,
		Color:       stageLabel.Color,
		Banner:      stageLabel.Banner,
	}
}

func PbStageToStage(stage *pb.PipelineStage) *Stage {
	return &Stage{
		ID:          stage.Id,
		Name:        stage.Name,
		Description: stage.Description,
		PipelineID:  stage.PipelineId,
		Index:       stage.Index,
	}
}

func convertToReference(stage Stage) *Stage {
	return &Stage{
		ID:             stage.ID,
		Name:           stage.Name,
		Description:    stage.Description,
		PipelineID:     stage.PipelineID,
		Index:          stage.Index,
		SerialNumber:   stage.SerialNumber,
		StageLabelID:   stage.StageLabelID,
		StageLabel:     stage.StageLabel,
		OrganisationID: stage.OrganisationID,
		CreatedBy:      stage.CreatedBy,
		UpdatedBy:      stage.UpdatedBy,
		UpdatedAt:      stage.UpdatedAt,
		CreatedAt:      stage.CreatedAt,
		BranchId:       stage.BranchId,
		DepartmentId:   stage.DepartmentId,
	}
}

func convertToReferences(stages []Stage) []*Stage {
	var referenceStages []*Stage
	for _, stage := range stages {
		referenceStages = append(referenceStages, convertToReference(stage))
	}
	return referenceStages
}

func StagesToPbStages(stages []*Stage) []*pb.PipelineStage {
	var pbStages []*pb.PipelineStage
	for _, stage := range stages {
		pbStages = append(pbStages, StageToPbStage(stage))
	}
	return pbStages
}

func PbStagesToStages(stages []*pb.PipelineStage) []Stage {
	var pbStages []Stage
	for _, stage := range stages {
		pbStages = append(pbStages, *PbStageToStage(stage))
	}
	return pbStages
}

func StageLabelToPbStageLabels(stageLabels []*StageLabel) []*pb.StageLabel {
	var pbStageLabels []*pb.StageLabel
	for _, stageLabel := range stageLabels {
		pbStageLabels = append(pbStageLabels, StageLabelToPbStageLabel(stageLabel))
	}
	return pbStageLabels
}

func PbStageLabelToStageLabels(stageLabels []*pb.StageLabel) []*StageLabel {
	var pbStageLabels []*StageLabel
	for _, stageLabel := range stageLabels {
		pbStageLabels = append(pbStageLabels, PbStageLabelToStageLabel(stageLabel))
	}
	return pbStageLabels
}
