package internal

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

func (*TaskuluApi) GetActivities(projectId string) string {
	return fmt.Sprintf("/api/v1/projects/%v/activities", projectId)
}

func (*TaskuluApi) GetAuthUrl(appKey string, sessionKey string) string {
	return fmt.Sprintf("?app_key=%v&session_key=%v", appKey, sessionKey)
}
