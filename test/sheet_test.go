package test

import (
	"github.com/stretchr/testify/assert"
	"taskulu/internal"
	"testing"
)

func TestFindSheet(t *testing.T) {
	projectId := "123456"
	taskId := "5d46b04456ad667202008c23"
	sheetService := internal.NewSheet(log, task)
	sheet := sheetService.FindSheetByTaskId(projectId, taskId)
	assert.Equal(t, sheet.Title, "فروغ")
}
