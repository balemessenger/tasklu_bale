package internal

import (
	"taskulu/pkg"
	"time"
)

type BaleIntegration struct {
	log        *pkg.Logger
	baleHook   *BaleHook
	activity   *ActivityService
	projectId  string
	sheetTitle string
}

func NewBaleIntegration(log *pkg.Logger, baleHook *BaleHook, activity *ActivityService, projectId string, sheetTitle string) *BaleIntegration {
	return &BaleIntegration{
		log:        log,
		baleHook:   baleHook,
		activity:   activity,
		projectId:  projectId,
		sheetTitle: sheetTitle,
	}
}

func (b *BaleIntegration) Run() {
	go b.run(b.projectId, b.sheetTitle)
}

func (b *BaleIntegration) run(projectId string, sheetName string) {
	for {
		b.SendLastActivity(projectId, sheetName)
		time.Sleep(time.Second)
	}
}

func (b *BaleIntegration) SendLastActivity(projectId string, sheetName string) string {
	msg := b.activity.GetLastActivity(projectId, sheetName)
	result, err := b.baleHook.Send(msg)
	if err != nil {
		b.log.Error("BaleHook error::", err)
	}
	return result
}
