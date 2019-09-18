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
	Remove  = "حذف"
)

type ActivityService struct {
	log          *pkg.Logger
	taskulu      *taskulu.Client
	sheetService *SheetService
	date         time.Time
	conditions   []string
	isFilter     bool
}

func NewActivity(log *pkg.Logger, taskulu *taskulu.Client, sheetService *SheetService, date time.Time, isFilter bool) *ActivityService {
	return &ActivityService{
		log:          log,
		taskulu:      taskulu,
		sheetService: sheetService,
		date:         date,
		conditions:   []string{Status, Receipt, Create, Remove},
		isFilter:     isFilter,
	}
}

func (b *ActivityService) GetLastActivity(projectId string, sheetName string) string {
	body, err := b.taskulu.GetActivities(projectId, 3)
	if err != nil {
		b.log.Error("Activity::", err)
		return ""
	}
	last := body.Data[0]
	t := time.Unix(int64(last.CreatedAt), 0)
	if len(last.Content.Keys) == 0 {
		b.log.Warnf("Empty content key for projectId: %v, sheetName: %v", projectId, sheetName)
		return ""
	}
	taskID := last.Content.Keys[0].Ids.TaskID
	if t.After(b.date) && b.filterActivity(&last, projectId, taskID, sheetName) {
		b.date = t
		if b.isFilter {
			return b.getActivityMessage(last.Content.Message, last.Content.Keys, projectId, taskID)
		}
		msg := b.getNonFilterActivity(last.Content.Message, last.Content.Keys, projectId, taskID)
		if last.Content.Digest != "" {
			msg = msg + "\\n_" + last.Content.Digest + "_"
		}
		return msg

	}
	return ""
}

func (b *ActivityService) filterActivity(body *model.ActivitiesData, projectID, taskID, sheetName string) bool {
	c1 := pkg.GetUtils().ReverseContainsString(b.conditions, body.Content.Message)
	sheet := b.sheetService.FindSheetByTaskId(projectID, taskID)
	c2 := sheet.Title == sheetName
	return !b.isFilter || (c1 && c2)
}

func (b *ActivityService) getActivityMessage(message string, keys []model.Keys, projectID, taskID string) string {
	size := len(keys)
	if strings.Contains(message, Status) {
		return fmt.Sprintf("تسک [%s](https://taskulu.com/a/project/%v/tasks/%v) از وضعیت %s به وضعیت %s تغییر کرد.", keys[0].Value, projectID, taskID, keys[size-2].Value, keys[size-1].Value)
	} else if strings.Contains(message, Receipt) {
		return fmt.Sprintf("سررسید کار [%s](https://taskulu.com/a/project/%v/tasks/%v) به %s تغییر پیدا کرد.", keys[0].Value, projectID, taskID, keys[1].Value)
	} else if strings.Contains(message, Create) {
		return fmt.Sprintf("کار [%s](https://taskulu.com/a/project/%v/tasks/%v) درلیست %s ایجاد شد.", keys[0].Value, projectID, taskID, keys[size-1].Value)
	} else {
		return ""
	}
}

func (b *ActivityService) getNonFilterActivity(message string, keys []model.Keys, projectID, taskID string) string {
	size := len(keys)
	msg := ""
	formattedMsg := fmt.Sprintf(
		"[%s](https://taskulu.com/a/project/%v/tasks/%v)",
		keys[0].Value,
		keys[0].Ids.ProjectID,
		keys[0].Ids.TaskID)
	if size == 2 {
		msg = strings.ReplaceAll(message, "%[0]", formattedMsg)
		msg = strings.ReplaceAll(msg, "%[1]", keys[1].Value.(string))
	} else if size == 3 {
		msg := strings.ReplaceAll(message, "%[0]", formattedMsg)
		msg = strings.ReplaceAll(msg, "%[1]", keys[1].Value.(string))
		msg = strings.ReplaceAll(msg, "%[2]", keys[2].Value.(string))
	} else if size == 4 {
		msg = strings.ReplaceAll(message, "%[0]", formattedMsg)
		msg = strings.ReplaceAll(msg, "%[1]", keys[1].Value.(string))
		msg = strings.ReplaceAll(msg, "%[2]", keys[2].Value.(string))
		msg = strings.ReplaceAll(msg, "%[3]", keys[3].Value.(string))
	}

	return msg
}
