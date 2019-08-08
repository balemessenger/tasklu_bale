package internal

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"taskulu/internal/model"
	"taskulu/pkg"
)

type TaskuluClient struct {
	log       *pkg.Logger
	baseUrl   string
	username  string
	password  string
	appKey    string
	sessionId string
}

type Option struct {
	BaseUrl  string
	Username string
	Password string
}

func NewTaskulu(log *pkg.Logger, option Option) *TaskuluClient {
	return &TaskuluClient{
		log:       log,
		baseUrl:   option.BaseUrl,
		username:  option.Username,
		password:  option.Password,
		appKey:    "",
		sessionId: "",
	}
}

func (t *TaskuluClient) CreateSession(username, password string) (error, *model.Session) {
	url := t.baseUrl + GetTaskuluApi().CreateSession()
	payload := fmt.Sprintf(`{"identifier":"%s","password":"%s"}`, username, password)

	resp, err := http.Post(url, "application/json", strings.NewReader(payload))

	if err != nil {
		return err, nil
	}
	if resp.StatusCode != 201 {
		return errors.New(resp.Status), nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}
	b := &model.Session{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return err, nil
	}
	return nil, b
}

func (t *TaskuluClient) GetActivities(projectId string) (error, *model.Activities) {
	resp, err := http.Get(t.getActivitiesUrl(projectId))
	if err != nil {
		return err, nil
	}
	if resp.StatusCode == 401 {

	}

	if resp.StatusCode != 200 {
		return errors.New(resp.Status), nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil
	}

	b := &model.Activities{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return err, nil
	}
	return nil, b
}

func (t *TaskuluClient) getActivitiesUrl(projectId string) string {
	return t.baseUrl + GetTaskuluApi().GetActivities(projectId) + GetTaskuluApi().GetAuthUrl(t.appKey, t.sessionId)

}
