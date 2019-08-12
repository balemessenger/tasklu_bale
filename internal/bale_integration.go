package internal

import (
	"taskulu/pkg"
	"time"
)

type BaleIntegration struct {
	log      *pkg.Logger
	baleHook *BaleHook
	activity ActivityService
}

func NewBaleIntegration(log *pkg.Logger) *BaleIntegration {
	return &BaleIntegration{
		log: log,
	}
}

func (b *BaleIntegration) Run() {
	go b.run("5a8d1fff56ad660b0dd0d343", "")
}

func (b *BaleIntegration) run(projectId string, sheetName string) {
	for {
		b.activity.SendLastActivity(projectId, sheetName)
		time.Sleep(time.Second)
	}
}
