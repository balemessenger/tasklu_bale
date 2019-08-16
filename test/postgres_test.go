package test

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	postgres2 "taskulu/internal/postgres"
	"testing"
)

func TestUpsertTaskulu(t *testing.T) {
	var tasks []postgres2.Taskulu
	err := postgres.UpsertTaskulu(11, "test", "test")
	assert.Nil(t, err)
	err = postgres.UpsertTaskulu(10, "test", "test2")
	tasks, err = postgres.GetTaskuluByUser(10)
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Username, "test")
	assert.Equal(t, tasks[0].Password, "test2")
}

func TestUpsertTaskuluByUsername(t *testing.T) {
	var tasks []postgres2.Taskulu
	userID := rand.Int31()
	err := postgres.UpsertTaskuluByUsername(int(userID), "test")
	assert.Nil(t, err)
	tasks, err = postgres.GetTaskuluByUser(int(userID))
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Username, "test")
	assert.Equal(t, tasks[0].Password, "")
}

func TestUpsertTaskuluByPassword(t *testing.T) {
	var tasks []postgres2.Taskulu
	userID := rand.Int31()
	err := postgres.UpsertTaskuluByPassword(int(userID), "pass")
	assert.Nil(t, err)
	tasks, err = postgres.GetTaskuluByUser(int(userID))
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)
	assert.Equal(t, tasks[0].Username, "")
	assert.Equal(t, tasks[0].Password, "pass")
}
