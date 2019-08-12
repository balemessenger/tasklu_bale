package internal

import (
	"fmt"
	"strings"
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"taskulu/pkg/taskulu/model"
	"time"
)

const (
	Status  = "وضعیت کار"
	Receipt = "سررسید"
	Create  = "را در لیست"
)

type ActivityService struct {
	log          *pkg.Logger
	taskulu      *taskulu.Client
	sheetService *SheetService
	date         time.Time
	conditions   []string
}

func NewActivity(log *pkg.Logger, taskulu *taskulu.Client, sheetService *SheetService, date time.Time) *ActivityService {
	return &ActivityService{
		log:          log,
		taskulu:      taskulu,
		sheetService: sheetService,
		date:         date,
		conditions:   []string{Status, Receipt, Create},
	}
}

func (b *ActivityService) GetLastActivity(projectId string, sheetName string) string {
	body, err := b.taskulu.GetActivities(projectId, 3)
	if err != nil {
		b.log.Error(err)
		return ""
	}
	last := body.Data[0]
	t := time.Unix(int64(last.CreatedAt), 0)
	if t.After(b.date) && b.filterActivity(&last, projectId, sheetName) {
		b.date = t
		return b.getActivityMessage(last.Content.Message, last.Content.Keys)
	}
	return ""
}

func (b *ActivityService) filterActivity(body *model.ActivitiesData, projectId string, sheetName string) bool {
	c1 := pkg.GetUtils().ReverseContainsString(b.conditions, body.Content.Message)
	taskId := body.Content.Keys[0].Ids.TaskID
	sheet := b.sheetService.FindSheetByTaskId(projectId, taskId)
	c2 := sheet.Title == sheetName
	return c1 && c2
}

func (b *ActivityService) getActivityMessage(message string, keys []model.Keys) string {
	if strings.Contains(message, Status) {
		return fmt.Sprintf("تسک %s از وضعیت %s به وضعیت %s تغییر کرد.", keys[2].Value, keys[1].Value, keys[0].Value)
	} else if strings.Contains(message, Receipt) {
		return fmt.Sprintf("سررسید کار %s به %s تغییر پیدا کرد.", keys[1].Value, keys[0].Value)
	} else if strings.Contains(message, Create) {
		return fmt.Sprintf("کار %s درلیست %s ایجاد شد.", keys[1].Value, keys[0].Value)
	} else {
		return ""
	}
}
