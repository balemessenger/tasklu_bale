package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"taskulu/internal"
	"taskulu/internal/model"
	"testing"
)

func TestResponse(t *testing.T) {
	input := `{
    "ok": true,
    "status": "OK",
    "data": [
        {
            "by": "56151bcafa1bc7a1810027ca",
            "content": {
                "keys": [
                    {
                        "type": "text",
                        "ids": {
                            "project_id": "5a8d1fff56ad660b0dd0d343",
                            "task_id": "5d46b04456ad667202008c23"
                        },
                        "value": "راه اندازی gitlab واسه CICD"
                    },
                    {
                        "type": "text",
                        "value": "Done"
                    },
                    {
                        "type": "text",
                        "value": "Doing"
                    }
                ],
                "message": "شما وضعیت کار %[0] را از %[1] به %[2] تغییر دادید."
            },
            "created_at": 1565091225
        }
	]
}`
	b := model.Body{}
	err := json.Unmarshal([]byte(input), &b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, b.Data[0].By, "56151bcafa1bc7a1810027ca")
	assert.Equal(t, b.Data[0].CreatedAt, 1565091225)

}

func TestRealResponse(t *testing.T) {
	taskulu := internal.NewTaskulu()
	err, body := taskulu.GetActivities("oRjJhqNBKGsSeZz5DOfbTAxCV_qAqalolbQMqLqisW7OmVyKvf5cxQYDpiSwGAePf5WBx74jH9IP09_QKxa4xhTDVTWurosLWpOK6VFGzIsRVLighsGEL_KOyXJe9on7", "77a426f5d99770633459fcb99dbb2975", "5a8d1fff56ad660b0dd0d343")

	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, body.Ok, true)
	assert.Equal(t, body.Status, "OK")

}
