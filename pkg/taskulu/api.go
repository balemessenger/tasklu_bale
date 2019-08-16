package taskulu

import (
	"fmt"
	"sync"
)

var (
	once sync.Once
	url  *TaskuluApi
)

type TaskuluApi struct{}

func GetTaskuluApi() *TaskuluApi {
	once.Do(func() {
		url = &TaskuluApi{}
	})
	return url
}

func (*TaskuluApi) CreateSession() string {
	return "/api/v1/sessions/password"
}

func (*TaskuluApi) GetNotifications() string {
	return "/api/v1/notifications"
}

func (*TaskuluApi) MarkNotificationSeen(id string) string {
	return fmt.Sprintf("/api/v1/notifications/%v/task", id)
}

func (*TaskuluApi) MarkAllNotificationSeen() string {
	return "/api/v1/notifications"
}

func (*TaskuluApi) GetProject(projectId string) string {
	return "/api/v1/projects/" + projectId
}

func (*TaskuluApi) GetActivities(projectId string) string {
	return fmt.Sprintf("/api/v1/projects/%v/activities", projectId)
}

func (*TaskuluApi) GetAuthUrl(appKey string, sessionKey string) string {
	return fmt.Sprintf("?app_key=%v&session_id=%v", appKey, sessionKey)
}
