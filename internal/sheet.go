package internal

import (
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"taskulu/pkg/taskulu/model"
)

type SheetService struct {
	log     *pkg.Logger
	taskulu *taskulu.Client
}

func NewSheet(log *pkg.Logger, taskulu *taskulu.Client) *SheetService {
	return &SheetService{
		log:     log,
		taskulu: taskulu,
	}
}

func (s *SheetService) FindSheetByTaskId(projectId string, taskId string) model.Sheets {
	projects, err := s.taskulu.GetProjects(projectId, 3)
	if err != nil {
		s.log.Error(err)
	}
	for _, sheet := range projects.Data.Sheets {
		for _, list := range sheet.TaskLists {
			if pkg.GetUtils().ContainsString(list.TaskOrder, taskId) {
				return sheet
			}
		}
	}

	return model.Sheets{}
}
