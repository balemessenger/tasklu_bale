package internal

import (
	"context"
	"fmt"
	"strings"
	"taskulu/internal/bot"
	"taskulu/pkg"
	"taskulu/pkg/taskulu/notification"
	"time"
)

type TaskuluSDK struct {
	log     *pkg.Logger
	service *NotificationService
	bot     *bot.TaskuluBot
	userId  int
}

func NewSDK(log *pkg.Logger, service *NotificationService, bot *bot.TaskuluBot, userId int) *TaskuluSDK {
	return &TaskuluSDK{
		log:     log,
		service: service,
		bot:     bot,
		userId:  userId,
	}
}

func (t *TaskuluSDK) RunNotification(ctx context.Context) {
	go t.runNotification(ctx)
}

func (t *TaskuluSDK) runNotification(ctx context.Context) {
	for {
		if ctx.Err() == context.Canceled {
			t.log.Debug("Taskulu::", "Cancelling the old goroutin  for user: ", t.userId)
			break
		}
		body, err := t.service.GetUnreadNotifications()
		if err != nil {
			t.log.Error("SDK::", err)
			time.Sleep(time.Second)
			continue
		}

		for _, b := range body {
			err := t.bot.SendTextMessage(t.userId, t.getMessage(b.Content))
			if err != nil {
				t.log.Error("SDK::", err)
			}
		}

		t.service.MarkAllSeen()

		time.Sleep(5 * time.Second)

	}
}

func (t *TaskuluSDK) getMessage(content notification.Content) string {
	msg := content.Message
	if len(content.Keys) == 3 {
		formattedMsg := fmt.Sprintf(
			"[%s](https://taskulu.com/a/project/%v/tasks/%v)",
			content.Keys[0].Value,
			content.Keys[0].Ids.ProjectID,
			content.Keys[0].Ids.TaskID)
		msg = strings.ReplaceAll(msg, "%[0]", formattedMsg)
		msg = strings.ReplaceAll(msg, "%[1]", content.Keys[1].Value)
		msg = strings.ReplaceAll(msg, "%[2]", content.Keys[2].Value)
	}

	return msg
}
