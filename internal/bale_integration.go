package internal

import (
	"taskulu/pkg"
	"time"
)

type BaleIntegration struct {
	log      *pkg.Logger
	baleHook *BaleHook
	activity *ActivityService
}

func NewBaleIntegration(log *pkg.Logger, baleHook *BaleHook, activity *ActivityService) *BaleIntegration {
	return &BaleIntegration{
		log:      log,
		baleHook: baleHook,
		activity: activity,
	}
}

func (b *BaleIntegration) Run() {
	go b.run("5a8d1fff56ad660b0dd0d343", "")
}

func (b *BaleIntegration) run(projectId string, sheetName string) {
	for {

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
