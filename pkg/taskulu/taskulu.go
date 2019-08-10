package taskulu

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"taskulu/pkg/taskulu/model"
	"taskulu/pkg"
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

func (t *Client) CreateSession(username, password string) (error, *model.Session) {
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

func (t *Client) GetActivities(projectId string, retryCount int) (error, *model.Activities) {
	resp, err := http.Get(t.getActivitiesUrl(projectId))
	if err != nil {
		return err, nil
	}

	// retry
	if resp.StatusCode == 401 {
		if retryCount == 0 {
			return errors.New(resp.Status), nil
		}
		t.retrySession()
		return t.GetActivities(projectId, retryCount-1)
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

func (t *Client) getActivitiesUrl(projectId string) string {
	return t.baseUrl + GetTaskuluApi().GetActivities(projectId) + GetTaskuluApi().GetAuthUrl(t.appKey, t.sessionId)

}

func (t *Client) retrySession() {
	err, s := t.CreateSession(t.username, t.password)
	t.logError(err)
	t.appKey = s.Data.AppKey
	t.sessionId = s.Data.SessionId
	t.retryCount++
}

func (t *Client) logError(err error) {
	t.log.Error(err)
}
