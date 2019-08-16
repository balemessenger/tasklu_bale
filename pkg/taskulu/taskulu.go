package taskulu

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"taskulu/pkg"
	"taskulu/pkg/taskulu/model"

	"github.com/pkg/errors"
	"taskulu/pkg/taskulu/notification"
)

type Client struct {
	log        *pkg.Logger
	baseUrl    string
	username   string
	password   string
	appKey     string
	sessionId  string
	retryCount int
}

type Option struct {
	BaseUrl  string
	Username string
	Password string
}

func New(log *pkg.Logger, option Option) *Client {
	return &Client{
		log:       log,
		baseUrl:   option.BaseUrl,
		username:  option.Username,
		password:  option.Password,
		appKey:    "",
		sessionId: "",
	}
}

func (t *Client) CreateSession(username, password string) (*model.Session, error) {
	url := t.baseUrl + GetTaskuluApi().CreateSession()
	payload := fmt.Sprintf(`{"identifier":"%s","password":"%s"}`, username, password)
	resp, err := http.Post(url, "application/json", strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 201 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	b := &model.Session{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *Client) GetNotifications(retryCount int) (*notification.Notifications, error) {
	resp, err := http.Get(t.getNotificationUrl())
	if err != nil {
		return nil, err
	}

	// retry
	if resp.StatusCode == 401 {
		if retryCount == 0 {
			return nil, errors.New(resp.Status)
		}
		t.retrySession()
		return t.GetNotifications(retryCount - 1)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	b := &notification.Notifications{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *Client) GetProjects(projectId string, retryCount int) (*model.Project, error) {
	resp, err := http.Get(t.getProjectsUrl(projectId))
	if err != nil {
		return nil, err
	}

	// retry
	if resp.StatusCode == 401 {
		if retryCount == 0 {
			return nil, errors.New(resp.Status)
		}
		t.retrySession()
		return t.GetProjects(projectId, retryCount-1)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	b := &model.Project{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *Client) GetActivities(projectId string, retryCount int) (*model.Activities, error) {
	url := t.getActivitiesUrl(projectId)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// retry
	if resp.StatusCode == 401 {
		if retryCount == 0 {
			return nil, errors.New(resp.Status)
		}
		t.retrySession()
		return t.GetActivities(projectId, retryCount-1)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	b := &model.Activities{}
	err = json.Unmarshal(body, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (t *Client) getActivitiesUrl(projectId string) string {
	return t.baseUrl + GetTaskuluApi().GetActivities(projectId) + GetTaskuluApi().GetAuthUrl(t.appKey, t.sessionId)
}

func (t *Client) getProjectsUrl(projectId string) string {
	return t.baseUrl + GetTaskuluApi().GetProject(projectId) + GetTaskuluApi().GetAuthUrl(t.appKey, t.sessionId)
}

func (t *Client) getNotificationUrl() string {
	return t.baseUrl + GetTaskuluApi().GetNotifications() + GetTaskuluApi().GetAuthUrl(t.appKey, t.sessionId)
}

func (t *Client) retrySession() {
	s, err := t.CreateSession(t.username, t.password)
	if err != nil {
		t.log.Error("Taskulu::", err)
		return
	}
	t.appKey = s.Data.AppKey
	t.sessionId = s.Data.SessionId
	t.retryCount++
}
