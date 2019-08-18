package internal

import (
	"context"
	"taskulu/internal/bot"
	"taskulu/pkg"
	"taskulu/pkg/taskulu"
	"time"
)

func RunIntegration(log *pkg.Logger, groupToken, projectId, sheetTitle string, isFilter bool) {
	bale := NewBale("https://api.bale.ai", groupToken)
	task := taskulu.New(log, taskulu.Option{
		BaseUrl:  "https://taskulu.com",
		Username: "bale",
		Password: "bale2513060",
	})
	sheet := NewSheet(log, task)
	activity := NewActivity(log, task, sheet, time.Now(), isFilter)
	integration := NewBaleIntegration(log, bale, activity, projectId, sheetTitle)
	integration.Run()
}

func RunNotification(ctx context.Context, log *pkg.Logger, bot *bot.TaskuluBot, userId int, username, password string) {
	client := taskulu.New(log, taskulu.Option{
		BaseUrl:  "https://taskulu.com",
		Username: username,
		Password: password,
	})
	service := NewNotification(log, client)
	sdk := NewSDK(log, service, bot, userId)
	sdk.RunNotification(ctx)
}
