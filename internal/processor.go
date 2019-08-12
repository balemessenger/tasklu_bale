package internal

import (
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"time"
)

func RunIntegration(log *pkg.Logger, groupToken, projectId, sheetTitle string) {
	bale := NewBale("https://api.bale.ai", groupToken)
	task := taskulu.New(log, taskulu.Option{
		BaseUrl:  "https://taskulu.com",
		Username: "amsjavan",
		Password: "ams2513060",
	})
	sheet := NewSheet(log, task)
	activity := NewActivity(log, task, sheet, time.Now())
	integration := NewBaleIntegration(log, bale, activity, projectId, sheetTitle)
	integration.Run()
}
