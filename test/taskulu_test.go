package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetActivites(t *testing.T) {
	projectId := "123456"

	b, err := task.GetActivities(projectId, 3)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, b.Data[0].By, "56151bcafa1bc7a1810027ca")
	assert.Equal(t, b.Data[0].CreatedAt, 1565091225)

}

func TestCreateSession(t *testing.T) {
	session, err := task.CreateSession("test", "test")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, session.Status, "CREATED")

}

func TestGetProjects(t *testing.T) {
	projectId := "123456"

	b, err := task.GetProjects(projectId, 3)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, b.Data.ID, "5b029ec356ad663c2600095d")
	assert.Equal(t, b.Data.Sheets[0].ID, "5a8d1fff56ad660b0dd0d345")
	assert.Contains(t, b.Data.Sheets[0].TaskLists[0].TaskOrder, "5d46b04456ad667202008c23")

}

func TestGetNotification(t *testing.T) {
	b, err := task.GetNotifications(3)

	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, b.Data.TotalUnseen, 4)
	assert.Equal(t, b.Data.Notifications[0].ID, "5b029ee056ad663c260013ee")
	assert.Contains(t, b.Data.Notifications[0].Content.Keys[0].Ids.ProjectID, "5b029ee056ad663c260013c9")

}
