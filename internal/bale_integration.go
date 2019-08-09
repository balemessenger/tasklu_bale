package internal

import (
	"fmt"
	"strings"
	"taskulu/internal/model"
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"time"
)

type BaleIntegration struct {
	log      *pkg.Logger
	taskulu  *taskulu.Client
	baleHook *BaleHook
	date     time.Time
}

func NewBaleIntegration(log *pkg.Logger, taskulu *taskulu.Client, baleHook *BaleHook, date time.Time) *BaleIntegration {
	return &BaleIntegration{
		log:      log,
		taskulu:  taskulu,
		baleHook: baleHook,
		date:     date,
	}
}

func (b *BaleIntegration) Run() {
	go b.run("5a8d1fff56ad660b0dd0d343")
}

func (b *BaleIntegration) run(projectId string) {
	for {
		b.SendLastActivity(projectId)
		time.Sleep(time.Second)
	}
}

func (b *BaleIntegration) SendLastActivity(projectId string) string {
	err, body := b.taskulu.GetActivities(projectId, 3)
	if err != nil {
		b.log.Error(err)
	}
	t := time.Unix(int64(body.Data[0].CreatedAt), 0)
	if t.After(b.date) {
		result, err := b.baleHook.Send(b.getActivityMessage(body.Data[0].Content.Message, body.Data[0].Content.Keys))
		if err != nil {
			b.log.Error("BaleHook error::", err)
		}
		b.date = t
		return result
	}
	return ""
}

func (b *BaleIntegration) getActivityMessage(message string, keys []model.Keys) string {
	if strings.Contains(message, "وضعیت کار") {
		return fmt.Sprintf("تسک %s از وضعیت %s به وضعیت %s تغییر کرد.", keys[2].Value, keys[1].Value, keys[0].Value)
	} else if strings.Contains(message, "سررسید") {
		return fmt.Sprintf("سررسید کار %s به %s تغییر پیدا کرد.", keys[1].Value, keys[0].Value)
	} else if strings.Contains(message, "را در لیست") {
		return fmt.Sprintf("کار %s درلیست %s ایجاد شد.", keys[1].Value, keys[0].Value)
	} else {
		return ""
	}
}
