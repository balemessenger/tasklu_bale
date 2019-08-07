package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"taskulu/internal/model"
)

type TaskuluApi struct {
	baseUrl string
}

func NewTaskulu(baseUrl string) *TaskuluApi {
	return &TaskuluApi{
		baseUrl: baseUrl,
	}
}

func (t *TaskuluApi) GetActivities(appKey string, sessionId string, projectId string) (error, *model.Body) {
	url := t.baseUrl + t.getActivitiesMethod(projectId) + "?app_key=" + appKey + "&session_id=" + sessionId
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	if resp.StatusCode != 200 {
		return errors.New(resp.Status), nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	b := &model.Body{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return err, nil
	}
	return nil, b
}

func (t *TaskuluApi) getActivitiesMethod(projectId string) string {
	return "/api/v1/projects/" + projectId + "/activities"
}
