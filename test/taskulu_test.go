package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetActivites(t *testing.T) {
	projectId := "123456"

	err, b := taskulu.GetActivities(projectId)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, b.Data[0].By, "56151bcafa1bc7a1810027ca")
	assert.Equal(t, b.Data[0].CreatedAt, 1565091225)

}

func TestCreateSession(t *testing.T) {
	err, session := taskulu.CreateSession("test", "test")

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, session.Status, "CREATED")

}
